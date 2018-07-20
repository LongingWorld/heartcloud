package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"xinyun/model"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func PostHandler(c *gin.Context) {
	fmt.Println("@@@@@@@Start@@@@@@@")
	/*连接数据库*/
	// db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxj2018?charset=utf8&parseTime=true&loc=Local")
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxjdata?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	/*关闭连接数据库*/
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
		return
	}

	fmt.Printf("*****@@@@@@   count is %d name is %s****\n", count, gauge.Name)
	if gauge.ID == 0 {
		log.Printf("量表%s不存在", gaugeID)
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

	/*读取redis获取token,通过token匹配用户信息BEGIN*/
	//获取Token
	authorization := c.GetHeader("Authorization")
	fmt.Printf("@@@@@@@   authorization is :%s\n", authorization)
	tokenKey := model.AccessTokenPrefix + authorization
	//连接Redis
	conRedis, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer conRedis.Close()
	clientInfo, err := redis.StringMap(conRedis.Do("Get", tokenKey))
	fmt.Printf("@@@@@@   客户端信息:%v\n", clientInfo)

	/*读取redis获取token,通过token匹配用户信息END*/

	/*获取xy_reportsetting报告设置表信息*/
	var reportSet model.Reportsetting
	if err := db.Table("xy_report_seating").
		Where("gauge_id = ?", gint).
		Find(&reportSet).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_report_setting error!", file, line)
		return
	}

	fmt.Printf("@@@@@@   reportSet is : \n %v\n", reportSet)

	var staffAnswer []model.StaffAnswer

	/*获取员工答题信息*/
	if err := db.Table("xy_staff_answer").
		Where("service_use_staff_id = ? AND staff_id = ? AND is_finish = ? AND deleted_at is null", staAnsID, 53, 2).
		Find(&staffAnswer).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff_answer error!", file, line)
		return
	}
	if len(staffAnswer) == 0 {
		log.Printf("员工%d答题信息表记录不存在！", 53)
		return
	}
	fmt.Printf("@@@@@@   员工ID：53,员工答题ID：%d ,答题信息:\n %v \n", staffAnswer[0].ID, staffAnswer)

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
	if err := tx.Table("xy_staff_answer").
		Where("service_use_staff_id = ? AND staff_id = ? AND is_finish = ? AND deleted_at is null", staAnsID, 53, 2).
		Updates(map[string]interface{}{"is_finish": 1, "deleted_at": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
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
			StaffID:         53,
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
		StaffID:             53,
		StaffName:           "宋志勇",
		StaffAge:            37,
		CompanyID:           12,
		CompanyName:         "四川心云智慧科技有限公司",
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
	if gauge.TemplateID == 4 {
		/*PHP逻辑$staffDimensionScore、$explainStaff，主要写入的即是这两张表*/
		/* 此为Golang逻辑,新建xy_subject_answer_map_explain(题目答案映射关系解释表),读出员工测试解释内容 */
		//var mapExplains = make([]model.MapExplain, 0)
		var staffDimension = make([]model.StaffDimension, 0)
		for key, value := range subjectsAnswersArr {
			var mapExplains = make([]model.ExplainsDetail, 0)
			var mapExplain model.ExplainsDetail
			var staffDim model.StaffDimension
			k, _ := strconv.Atoi(key)
			fmt.Printf("@@@@@@   subject_id is :%d，subject_answer_id is:%d \n", k, value)
			if k == 354 || k == 355 || k == 356 || k == 357 {
				if err := tx.Debug().Raw(`SELECT 
					a.map_id,b.option_name,a.map_explain,a.map_proposal
				FROM
					xy_subject_answer_map_explain a,
					xy_subject_answer b
				WHERE
					a.subject_answer_id = b.id
					AND a.subject_id = ?
						AND a.subject_answer_id = ?
						AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
					_, file, line, _ := runtime.Caller(0)
					log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
					continue
					//回滚事务
					//tx.Rollback()
					//return
				}
				/* if err := db.Table("xy_subject_answer_map_explain").
					Select("map_id", "map_name", "map_explain", "map_proposal").
					Where("subject_id = ?", k).
					Where("subject_answer_id = ? AND map_sqe_no = ?", value, 1).
					Scan(&mapExplain).Error; err != nil {
					_, file, line, _ := runtime.Caller(0)
					log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
					//回滚事务
					tx.Rollback()
					return
				} */
			} else if k == 358 {
				/*第五题:面对上级，总是忐忑不安，内心紧张是高焦虑的一种表现；见到上级躲着走是一种高回避的表现。
				根据你和上级大多数在一起的情况，你认为自己属于？*/
				if value == 1130 { /*高焦虑，低回避*/
					if (subjectsAnswersArr["357"] == 1123 || subjectsAnswersArr["357"] == 1124) &&
						(subjectsAnswersArr["356"] == 1119 || subjectsAnswersArr["356"] == 1120 ||
							subjectsAnswersArr["356"] == 1121 || subjectsAnswersArr["356"] == 1122) {
						if err := tx.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 2).Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
							continue
							//回滚事务
							//tx.Rollback()
							//return
						}
						/* if err := db.Table("xy_subject_answer_map_explain").
							Select("map_id", "map_name", "map_explain", "map_proposal").
							Where("subject_id = ?", k).
							Where("subject_answer_id = ? AND map_sqe_no = ?", value, 2).
							Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
							//回滚事务
							tx.Rollback()
							return
						} */
					} else {
						if err := tx.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
							continue
							//回滚事务
							//tx.Rollback()
							// return
						}
						/* if err := db.Table("xy_subject_answer_map_explain").
							Select("map_id", "map_name", "map_explain", "map_proposal").
							Where("subject_id = ?", k).
							Where("subject_answer_id = ? AND map_sqe_no = ?", value, 1).
							Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
							//回滚事务
							tx.Rollback()
							return
						} */
					}
				} else if value == 1132 { /*低焦虑，低回避  */
					if (subjectsAnswersArr["356"] == 1119 || subjectsAnswersArr["356"] == 1120) &&
						(subjectsAnswersArr["357"] == 1123 || subjectsAnswersArr["357"] == 1124) {
						if err := tx.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
							continue
							//回滚事务
							// tx.Rollback()
							// return
						}
						/* if err := db.Table("xy_subject_answer_map_explain").
							Select("map_id", "map_name", "map_explain", "map_proposal").
							Where("subject_id = ?", k).
							Where("subject_answer_id = ? AND map_sqe_no = ?", value, 1).
							Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
							//回滚事务
							tx.Rollback()
							return
						} */
					} else /* if (subjectsAnswersArr["356"] == 1119 || subjectsAnswersArr["356"] == 1120 ||
					subjectsAnswersArr["356"] == 1121 || subjectsAnswersArr["356"] == 1122) ||
					(subjectsAnswersArr["357"] == 1125 || subjectsAnswersArr["357"] == 1126 ||
						subjectsAnswersArr["357"] == 1127 || subjectsAnswersArr["357"] == 1128) */{
						if err := tx.Raw(`SELECT 
							a.map_id,b.option_name,a.map_explain,a.map_proposal
						FROM
							xy_subject_answer_map_explain a,
							xy_subject_answer b
						WHERE
							a.subject_answer_id = b.id
							AND a.subject_id = ?
								AND a.subject_answer_id = ?
								AND a.map_sqe_no = ?`, k, value, 2).Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
							continue
							//回滚事务
							// tx.Rollback()
							// return
						}
						/* if err := db.Table("xy_subject_answer_map_explain").
							Select("map_id", "map_name", "map_explain", "map_proposal").
							Where("subject_id = ?", k).
							Where("subject_answer_id = ? AND map_sqe_no = ?", value, 2).
							Scan(&mapExplain).Error; err != nil {
							_, file, line, _ := runtime.Caller(0)
							log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
							//回滚事务
							tx.Rollback()
							return
						} */
					}
				} else if value == 1129 || value == 1131 {
					if err := tx.Debug().Raw(`SELECT 
						a.map_id,b.option_name,a.map_explain,a.map_proposal
					FROM
						xy_subject_answer_map_explain a,
						xy_subject_answer b
					WHERE
						a.subject_answer_id = b.id
						AND a.subject_id = ?
							AND a.subject_answer_id = ?
							AND a.map_sqe_no = ?`, k, value, 1).Scan(&mapExplain).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:%s:Select Table xy_subject_answer_map_explain error!", file, line, err)
						continue
						//回滚事务
						//tx.Rollback()
						//return
					}
				}
				//打印注释
				fmt.Printf("@@@@@@   subjectsAnswersArr[356] is %d\n subjectsAnswersArr[357] is %d\n",
					subjectsAnswersArr["356"], subjectsAnswersArr["357"])
				var ParaConf model.SystemConf
				if subjectsAnswersArr["356"] == 1119 && subjectsAnswersArr["357"] == 1125 {
					if err := tx.Debug().Table("xy_system_conf").
						Select("para_conf").
						Where("para_id = ?", "XYXJLB4").
						Scan(&ParaConf).Error; err != nil {
						_, file, line, _ := runtime.Caller(0)
						log.Printf("%s:%d:Select Table xy_system_conf error!", file, line)
						continue
						//回滚事务
						// tx.Rollback()
						// return
					}
					/*拼接额外头部信息*/
					fmt.Printf("@@@@@@   ParaConf is : \n %v \n", ParaConf)
					midStr := []byte(ParaConf.ParaConf)
					mapEx := []byte(mapExplain.MapExplain)
					midStr = append(midStr, mapEx...)
					mapExplain.MapExplain = string(midStr)
					fmt.Printf("@@@@@@   midStr is : \n %v \n", string(midStr))
					fmt.Printf("@@@@@@   mapExplain is : \n %v \n", mapExplain.MapExplain)
				}
			}
			mapExplains = append(mapExplains, mapExplain)
			staffDim.ExplainsDetail = mapExplains

			/*查询dim_name维度名称*/
			type resultDim struct {
				DimensionID int
				DimName     string
				DimSuggest  string
				DimDesc     string
			}
			var res resultDim
			if err := tx.Table("xy_dim_and_subject a").
				Select("a.dimension_id,b.dim_name,b.dim_suggest,b.dim_desc").Joins("left join xy_dimension b on a.dimension_id = b.id").
				Where("a.subject_id = ?", k).Scan(&res).Error; err != nil {
				_, file, line, _ := runtime.Caller(0)
				log.Printf("%s:%d:Select Table xy_subject_answer_map_explain error!", file, line)
				continue
			}
			//fmt.Printf("@@@@@@   res is %v\n", res)

			staffDim.StaffID = 53
			staffDim.DimName = res.DimName
			//staffDim.NormName = "不知道什么原因"
			staffDim.DimensionID = res.DimensionID
			staffDim.DimSuggest = res.DimSuggest
			staffDim.DimDesc = res.DimDesc
			staffDimension = append(staffDimension, staffDim)
		}
		/*转换JSON*/
		staDimJSON, err := json.Marshal(staffDimension)
		if err != nil {
			log.Println("Marshal JSON字符串错！")
			//回滚事务
			tx.Rollback()
			return
		}
		// fmt.Printf("staDimJSON is %v\n", staDimJSON)

		reportStaffData.ReportData = string(staDimJSON)
		fmt.Printf("reportStaffData.ReportData is %s\n", reportStaffData.ReportData)
		// reportStaffData.ReportStaffID = int(reportStaff.ID)
		reportStaffData.ReportStaffID = int(reportStaff.ID)
	} else {
		log.Printf("%d未知的模板报告！", gauge.TemplateID)
	}

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
	if err := tx.Table("xy_report_staff_data").Create(&reportStaffData).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Insert Table xy_report_staff_data error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}

	/*更新员工报告report_staff答题得分*/
	if err := tx.Table("xy_report_staff").
		Where("id = ?", reportStaff.ID).
		Update("total_score", totalScore[0]).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Insert Table xy_subject_answer AND xy_staff_auswer_option error!", file, line)
		//回滚事务
		tx.Rollback()
		return
	}
	/*更新员工答题表xy_staff_answer得分*/
	if err := tx.Table("xy_staff_answer").
		Where("id = ?", staffAnsID).
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

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}
