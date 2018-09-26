package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"heartcloud/database"
	"heartcloud/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql驱动
)

/*GenerateStaffReport function generate staff reports*/
func GenerateStaffReport(c *gin.Context) {
	fmt.Println("@@@@@@@Start@@@@@@@")

	/* //验证登录Token信息，并获取用户信息
	staffInfo, err := verifyToken(c)
	if err != nil {
		log.Printf("验证Token信息失败！\n")
		c.JSON(http.StatusBadGateway, "Token失效")
		return
	} */
	//获取客户端用户信息
	staffInfo := c.Keys
	staffID := int(staffInfo["staff_id"].(float64))
	/* if !ok {
		fmt.Println("errors what!?")
	} */
	staffName := staffInfo["name"].(string)
	companyID := int(staffInfo["company_id"].(float64))
	companyName := staffInfo["company_name"].(string)

	/*连接数据库*/
	// db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxj2018?charset=utf8&parseTime=true&loc=Local")
	// db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxjdata?charset=utf8&parseTime=true&loc=Local")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// // 全局禁用表名复数
	// db.SingularTable(true)
	// /*关闭连接数据库*/
	db := database.ConnectDataBase()
	defer db.Close()

	var gauge model.Gauge
	// var company Company
	//var reportSet model.Reportsetting

	/*获取量表ID*/
	gaugeID := c.PostForm("gauge_id")
	//db.First(&gauge, "id = ?", gaugeID)
	gint, _ := strconv.Atoi(gaugeID)
	fmt.Printf("*****gauge_id is %d****\n", gint)
	//db.Table("xy_gauge").Where("id = ?", gint).Find(&xy_gauge)(
	if err := db.Table("xy_gauge").
		Where("id = ?", gint).
		Scan(&gauge).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_gauge error!", file, line)
		c.JSON(401, "系统异常")
		return
	}
	fmt.Printf("@@@@@@   xy_gauge量表数据:\n %v\n", gauge)
	var count int
	//db.Model(&model.Gauge{}).Where("id = ?", gint).Count(&count) //获取记录数
	if err := db.Table("xy_gauge").
		Select("id").Where("id = ?", gint).
		Count(&count).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_gauge error!", file, line)
		c.JSON(401, "系统异常")
		return
	}

	fmt.Printf("*****@@@@@@   count is %d name is %s****\n", count, gauge.Name)
	if gauge.ID == 0 {
		log.Printf("量表%s不存在", gaugeID)
		c.JSON(401, "量表不存在")
		return
	}

	/*获取员工答题ID*/
	staffAnsID := c.PostForm("service_use_staff_id")
	staAnsID, _ := strconv.Atoi(staffAnsID)
	fmt.Printf("@@@@@@   staffAnsID is %s,stafAnsID is %d \n", staffAnsID, staAnsID)

	/*获取答题答案JSON字符串*/
	subjectsAnswers := c.PostForm("subjects_answers")
	fmt.Printf("@@@@@@@   subjectsAnswers is :\n %s\n", subjectsAnswers)

	/*解析JSON字符串存入map中*/
	var subjectsAnswersArr = make(map[string]int)
	errs := json.Unmarshal([]byte(subjectsAnswers), &subjectsAnswersArr)
	if errs != nil {
		log.Println("解析JSON字符串错！")
		return
	}
	if len(subjectsAnswersArr) == 0 {
		log.Println("请答题！")
		return
	}
	fmt.Printf("@@@@@@   答案信息map:\n %v\n", subjectsAnswersArr)

	/*获取xy_reportsetting报告设置表信息*/
	var reportSet model.Reportsetting
	if err := db.Table("xy_report_setting").
		Where("gauge_id = ?", gint).
		Find(&reportSet).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_report_setting error!", file, line)
		return
	}

	fmt.Printf("@@@@@@   reportSet is : \n %v\n", reportSet)

	var staffAnswer []model.StaffAnswer

	/*获取员工答题信息*/
	if err := db.Debug().Table("xy_staff_answer").
		Where("service_use_staff_id = ? AND staff_id = ? AND is_finish = ? AND deleted_at is null", staAnsID, staffID, 2).
		Find(&staffAnswer).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff_answer error!", file, line)
		return
	}
	if len(staffAnswer) == 0 {
		log.Printf("员工%d答题信息表记录不存在！", staffID)
		return
	}
	fmt.Printf("@@@@@@   员工ID：%d,员工答题ID：%d ,答题信息:\n %v \n", staffID, staffAnswer[0].ID, staffAnswer)

	/*获取员工答题题目列表*/
	subjectIDs := getKeysFromMap(subjectsAnswersArr)
	if len(subjectIDs) == 0 {
		log.Println("请答完所有题目！")
		return
	}
	fmt.Printf("@@@@@@   员工答题题目列表 数目:%d \n %v\n", len(subjectIDs), subjectIDs)

	/*将取出的题目列表subjectIDs与量表对应的题目列表做比较，量表对应的所有题目存在于subjectIDs中*/
	var subCount int
	if err := db.Table("xy_subject").Where("gauge_id = ?", gaugeID).
		Not("id", subjectIDs).
		Count(&subCount).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_subject error!", file, line)
		return
	}
	fmt.Printf("@@@@@@   获取的题目列表是否与量表对应的题目列表相同，差异数目： %d\n", subCount)
	if subCount > 0 {
		log.Println("请答完所有题目！")
		return
	}
	fmt.Printf("@@@@@@   获取的题目列表是否与量表对应的题目列表相同，差异数目： %d\n", subCount)

	/*开启事务*/
	tx := db.Begin()
	/*更新staffanswer表isfinish和答题结束时间*/
	if err := tx.Debug().Table("xy_staff_answer").
		Where("service_use_staff_id = ? AND staff_id = ? AND is_finish = ? AND deleted_at is null",
			staAnsID, staffID, 2).
		Updates(map[string]interface{}{"is_finish": 1, "deleted_at": time.Now().Format("2006-01-02 15:04:05")}).
		Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Update Table xy_staff_answer error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}

	/*写入xy_staff_auswer_option员工答案信息表*/
	for key, value := range subjectsAnswersArr {
		k, _ := strconv.Atoi(key)
		staffAnsOpe := model.StaffAnswerOpetion{
			SubjectID:       k,
			SubjectAnswerID: value,
			StaffAnswerID:   int(staffAnswer[0].ID),
			StaffID:         staffID,
		}
		if err := tx.Table("xy_staff_auswer_option").Create(&staffAnsOpe).Error; err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:Insert Table xy_staff_auswer_option error!", file, line)
			//回滚事务
			tx.Rollback()
			return
		}
	}

	/*查询员工报告表最大记录ID*/
	var maxRepID []int
	if err := tx.Table("xy_report_staff").Select("ID").Pluck("MAX(ID)", &maxRepID).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_report_staff error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}
	maxRepID[0]++
	//计算员工年龄
	thisYear := time.Now().Year()
	var year []int
	if err := tx.Debug().Table("xy_staff").
		// Select("DATE_FORMAT(birthday,'%Y')").
		Where("id = ?", staffID).
		Pluck("DATE_FORMAT(birthday,'%Y')", &year).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}
	age := thisYear - year[0]
	fmt.Printf("######   This year is %d,birthday is %v,age is %d\n", thisYear, year, age)
	/*生成员工报告*/
	var reportStaff model.ReportStaff
	reportStaff = model.ReportStaff{
		StaffAnswerID:       staAnsID,
		Name:                gauge.ShowName,
		HideReportIntroduce: reportSet.HideReportIntroduce,
		Introduce:           reportSet.ReportIntroduce,
		HideDescribe:        reportSet.HideDescribe,
		Describe:            reportSet.Describe,
		HideShowMethod:      reportSet.HideShowMethod,
		ShowMethod:          reportSet.ShowMethod,
		HideCliches:         reportSet.HideCliches,
		Cliches:             reportSet.Cliches,
		HideComment:         reportSet.HideComment,
		Comment:             reportSet.Comment,
		HideDimSuggest:      reportSet.HideDimSuggest,
		GaugeID:             int(gauge.ID),
		StaffID:             staffID,
		StaffName:           staffName,
		StaffAge:            age,
		CompanyID:           companyID,
		CompanyName:         companyName,
		Number:              createRepNO("SNO"),
		Status:              1,
		TemplateID:          gauge.TemplateID,
		GenerateDate:        time.Now().Format("2006-01-02 15:04:05"),
		// Position       string
		// Marriage       int
		// TotalScore     float32
		// CreatedTime    string
		// UpdatedTime    string
		// DeletedTime    string
	}
	reportStaff.ID = uint(maxRepID[0])
	// fmt.Printf("@@@@@@   生成员工报告信息:\n %v\n", reportStaff)

	/*存入员工报告表*/
	if err := tx.Table("xy_report_staff").Create(&reportStaff).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Insert Table xy_report_staff error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}

	/*根据模板报告ID TemplateID，计算对应维度得分###新建流程，抛弃原先的流程###*/
	// if xy_gauge.TemplateID == 4 {
	// 	/*理清$staffDimensionScore、$explainStaff，主要写入的即是这两张表*/

	// } else {
	// 	log.Printf("%d未知的模板报告！", xy_gauge.TemplateID)
	// }

	var reportStaffData model.ReportStaffData //定义xy_report_staff_data数据库对应结构体结构体，写入数据库
	//定义通用数据结构类型
	var reportDataDtl interface{}
	var err error
	if gauge.TemplateID == 4 {
		/*PHP逻辑$staffDimensionScore、$explainStaff，主要写入的即是这两张表*/
		/* 此为Golang逻辑,新建xy_subject_answer_map_explain(题目答案映射关系解释表),读出员工测试解释内容 */
		//var mapExplains = make([]model.MapExplain, 0)
		reportDataDtl, err = GenerateStaffReportOfAuthorityRelation(tx, subjectsAnswersArr, staffAnswer[0])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:生成权威关系报告失败 error!", file, line)
			//回滚事务
			tx.Rollback()
		}
		/*转换JSON*/
		// staDimJSON, err := json.Marshal(staffDimension)
		// if err != nil {
		// 	log.Println("Marshal JSON字符串错！")
		// 	//回滚事务
		// 	tx.Rollback()
		// 	return
		// }
		// fmt.Printf("staDimJSON is %v\n", staDimJSON)

		// reportStaffData.ReportData = string(staDimJSON)
		// fmt.Printf("reportStaffData.ReportData is %s\n", reportStaffData.ReportData)
		// reportStaffData.ReportStaffID = int(reportStaff.ID)
	} else if gauge.TemplateID == 5 {
		reportDataDtl, err = GenerateStaffReportOfChronicFatigues(tx, subjectsAnswersArr)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:生成慢性疲劳报告失败 error!", file, line)
			//回滚事务
			tx.Rollback()
		}
		// fmt.Printf("######    chronicFatigueDetails \n  %v\n", chronicFatigueDetails)
		/*转换JSON*/
		// chroFatiDetail, err := json.Marshal(chronicFatigueDetails)
		// if err != nil {
		// 	log.Println("Marshal JSON字符串错！")
		// 	//回滚事务
		// 	tx.Rollback()
		// 	return
		// }
		// reportStaffData.ReportData = string(chroFatiDetail)
		// fmt.Printf("reportStaffData.ReportData is %s\n", reportStaffData.ReportData)
	} else if gauge.TemplateID == 6 {
		reportDataDtl, err = GenerateStaffReportOfEgoState(tx, subjectsAnswersArr, staffAnswer[0])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:生成自我状态报告失败 error!", file, line)
			//回滚事务
			tx.Rollback()
		}
		// fmt.Printf("######    EgoStateDetails \n  %v\n", egoStateDetails)
		// /*转换JSON*/
		// egoStateDetail, err := json.Marshal(egoStateDetails)
		// if err != nil {
		// 	log.Println("Marshal JSON字符串错！")
		// 	//回滚事务
		// 	tx.Rollback()
		// 	return
		// }
		// reportStaffData.ReportData = string(egoStateDetail)
		// fmt.Printf("reportStaffData.ReportData is %s\n", reportStaffData.ReportData)
	} else if gauge.TemplateID == 7 {
		reportDataDtl, err = GenerateStaffReportOfDSQ(tx, subjectsAnswersArr, staffAnswer[0])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Printf("%s:%d:生成防御方式报告失败 error!", file, line)
			//回滚事务
			tx.Rollback()
		}
		// fmt.Printf("######    DSQDetails \n  %v\n", DSQDetails)
		// /*转换JSON*/
		// DSQDetail, err := json.Marshal(DSQDetails)
		// if err != nil {
		// 	log.Println("Marshal JSON字符串错！")
		// 	//回滚事务
		// 	tx.Rollback()
		// 	return
		// }
		// reportStaffData.ReportData = string(DSQDetail)
		// fmt.Printf("reportStaffData.ReportData is %s\n", reportStaffData.ReportData)
	} else {
		log.Printf("%d未知的模板报告！", gauge.TemplateID)
	}

	//转换JSON格式
	reportDataDtlJSON, errs := json.Marshal(reportDataDtl)
	if errs != nil {
		log.Println("Marshal JSON字符串错！")
		//回滚事务
		tx.Rollback()
		return
	}
	reportStaffData.ReportData = string(reportDataDtlJSON)
	fmt.Printf("reportStaffData.ReportData is %s\n", reportStaffData.ReportData)

	reportStaffData.ReportStaffID = int(reportStaff.ID)

	/*计算员工报告中答题得分*/
	var totalScore []int
	/* db.Table("xy_subject_answer a").Select("count(*)").
	Joins("left join xy_staff_auswer_option b on (a.id = b.subject_answer_id)").
	Where("b.staff_answer_id = ?", int(staffAnswer[0].ID)).Count(&totalScore) */
	if err := tx.Table("xy_subject_answer a"). /* Select("count(*)"). */ /* "a.id,a.subject_id,a.fraction" */
							Joins("left join xy_staff_auswer_option b on (a.id = b.subject_answer_id)").
							Where("b.staff_answer_id = ?", int(staffAnswer[0].ID)).
							Pluck("SUM(a.fraction)", &totalScore).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_subject_answer AND xy_staff_auswer_option error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}
	//db.Raw("SELECT a.id,a.subject_id,a.fraction FROM xyxjdata.xy_subject_answer a LEFT JOIN xyxjdata.xy_staff_auswer_option b ON 	(a.id = b.subject_answer_id) WHERE b.staff_answer_id = ?", int(staffAnswer[0].ID)).Scan(&res)

	fmt.Printf("@@@@@@   员工报告ID:%d,员工答题总分: %d %v\n", reportStaff.ID, int(staffAnswer[0].ID), totalScore)

	/*生成员工报告数据表reportStaffData数据
	查询出xy_norm_explain中的f.id(常模说明项ID),
	f.name(说明项名称),f.score_introduce(分值说明),f.coach_proposal(辅导建议)
	然后以json格式写入reportStaffData->report_data_extra以及reportStaffData->report_data 中
	*/
	if err := tx.Debug().Table("xy_report_staff_data").Create(&reportStaffData).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Insert Table xy_report_staff_data error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}

	/*更新员工报告report_staff答题得分*/
	if err := tx.Debug().Table("xy_report_staff").
		Where("id = ?", reportStaff.ID).
		Update("total_score", totalScore[0]).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Insert Table xy_subject_answer AND xy_staff_auswer_option error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}
	/*更新员工答题表xy_staff_answer得分*/
	if err := tx.Debug().Table("xy_staff_answer").
		Where("service_use_staff_id = ?", staffAnsID).
		Update("score", totalScore[0]).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Insert Table xy_subject_answer AND xy_staff_auswer_option error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}

	// //若返回json数据，可以直接使用gin封装好的JSON方法

	//提交事务
	tx.Commit()
	// tx.Rollback()
	c.JSON(http.StatusOK, "success")
	return
}

/*获取map中的key值，并存放在slice中*/
func getKeysFromMap(m map[string]int) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func createRepNO(s string) string {
	result := []byte(s)
	timeSub := time.Now().Format("20060102150405")
	fmt.Printf("@@@@@@   timesub is %s\n", timeSub)
	now := []byte(timeSub)
	randstr := []byte(strconv.Itoa(GenerateRangeNum(1000, 9999)))
	fmt.Printf("@@@@@@   randstr is %s\n", string(randstr))
	result = append(result, now[6:10]...)
	result = append(result, randstr...)
	fmt.Printf("@@@@@@   SNO is %s\n", string(result))
	return string(result)
}

/* func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
} */

/*GenerateRangeNum generate range numbers*/
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}
