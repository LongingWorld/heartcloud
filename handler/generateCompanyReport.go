package handler

import (
	"encoding/json"
	"fmt"
	"heartcloud/model"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

/*GenerateCompanyReport function generate company reports*/
func GenerateCompanyReport(c *gin.Context) {
	fmt.Println("@@@@@@@GenerateCompanyReport()Begin@@@@@@@")
	//验证登录Token信息，并获取用户信息
	companyInfo, err := verifyToken(c)
	if err != nil {
		log.Printf("验证Token信息失败！\n")
		c.JSON(40001, "Token失效")
		return
	}
	fmt.Println(companyInfo)
	/* companyName, ok := companyInfo["name"].(string)
	if !ok {
		fmt.Println("errors what!?")
	}
	userName := companyInfo["username"].(string)
	companyID := int(companyInfo["company_id"].(float64))
	phone := companyInfo["phone"].(string)
	fmt.Printf("@@@@@@   companyName is :%s,userName is :%s,phone is :%s,company_id is :%d \n",
		companyName, userName, phone, companyID) */
	/*获取企业报告名称*/
	reportName := c.PostForm("report_name")
	/*获取企业ID*/
	compID := c.PostForm("company_id")
	companyID, _ := strconv.Atoi(compID)
	/*获取企业分发次数*/
	distriTime := c.PostForm("times")
	distributeTime, _ := strconv.Atoi(distriTime)
	/*获取量表列表JSON字符串*/
	gaugeLists := c.PostFormArray("gauge_lists")

	fmt.Printf("@@@@@@   reportName is %s,companyID is %d, distributeTime is %d,gaugeLists is %v\n",
		reportName, companyID, distributeTime, gaugeLists)
	if len(gaugeLists) == 0 {
		log.Println("请选择量表！")
		return
	}

	//转换类型string to int
	var gaugeIDs = make([]int, 0)
	for _, value := range gaugeLists {

		gaugeID, _ := strconv.Atoi(value)
		gaugeIDs = append(gaugeIDs, gaugeID)
		fmt.Printf("@@@@@@  value is %s,gaugeID is %d \n", value, gaugeID)
	}
	fmt.Printf("@@@@@@   gaugeIDs is %v\n", gaugeIDs)
	db := ConnectDataBase()
	/*关闭连接数据库*/
	defer db.Close()

	//获取量表列表信息
	var gauges []model.Gauge
	if err := db.Debug().Table("xy_gauge").
		Where("id in (?)", gaugeIDs).
		Scan(&gauges).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_gauge error!", file, line)
		c.JSON(401, "系统异常")
		return
	}

	/* 获取report_company_id插入的企业报告ID */
	/*获取企业报告ID*/
	repCompID := c.PostForm("report_company_id")
	reportCompanyID, _ := strconv.Atoi(repCompID)
	//企业报告report_company_data表中的report_data以及 report_data_api字段
	type ReportData struct {
		RepDataStr    string
		RepDataAPIStr string
	}
	var (
		reportDataStr ReportData
		repData       map[string]interface{}
		repDataAPI    map[string]interface{}
	)

	if err := db.Debug().Table("xy_report_company_data").
		Where("id = ?", reportCompanyID).
		Scan(&reportDataStr).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_gauge error!", file, line)
		c.JSON(401, "系统异常")
		return
	}

	//解析report_company_data.report_data to map
	if err := json.Unmarshal([]byte(reportDataStr.RepDataStr), &repData); err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:解析JSON字符串出错", file, line)
		c.JSON(401, "系统异常")
		return
	}
	fmt.Printf("@@@@@@Original   reportdata is: \n%v\n", repData)
	//解析report_company_data.report_data_api to map
	if err := json.Unmarshal([]byte(reportDataStr.RepDataAPIStr), &repDataAPI); err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:解析JSON字符串出错", file, line)
		c.JSON(401, "系统异常")
		return
	}
	fmt.Printf("@@@@@@Original   reportdataapi is: \n%v\n", repDataAPI)

	for _, gaugeinfo := range gauges {
		if gaugeinfo.TemplateID == 1 || gaugeinfo.TemplateID == 2 || gaugeinfo.TemplateID == 3 {
			continue
		} else if gaugeinfo.TemplateID == 4 {
			repComData, err := createCompanyReportData(db, gaugeinfo, companyID, distributeTime)
			if err != nil {
				c.JSON(401, "系统异常")
			}
			repData["template4"] = repComData
		} else {
			log.Println("量表ID无效！")
			return
		}
	}

	//map to json
	reportCompanyData, err := json.Marshal(&repData)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:解析JSON字符串出错", file, line)
		c.JSON(401, "系统异常")
		return
	}
	fmt.Printf("@@@@@@   new report_data is:\n %v\n", reportCompanyData)
	//更新xy_report_company_data.report_data数据
	if err := db.Debug().Table("xy_report_company_data").
		Update("report_data", string(reportCompanyData)).
		Where("id = ?", reportCompanyID).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Update Table xy_report_company_data error!", file, line)
		c.JSON(401, "系统异常")
		return
	}

	c.JSON(http.StatusOK, "success")
	fmt.Println("@@@@@@@GenerateCompanyReport()end@@@@@@@")
	return
}

func createCompanyReportData(db *gorm.DB, gauge model.Gauge, comID int, comTimes int) (map[string]interface{}, error) {
	var template map[string]interface{}
	template["id"] = gauge.ID
	template["name"] = gauge.Name
	template["show_name"] = gauge.ShowName
	template["template_id"] = gauge.TemplateID
	template["section1"] = `<p style="font-size:16px;line-height:2em;text-indent:2em;">“人际关系之权威关系投射测验”旨在反应街道当前管理效力的现状，引导员工更好的建设上下级关系，提升团队凝聚力和执行力。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">人际关系之权威关系投射测验属于心理学中的投射测验（Projective Test），它是以心理动力学理论为基础，以多种“动物”形象探索被试当前的权威关系模式，从而引导被试构建更加高效、良性互动的权威关系。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">“权威关系”在生活中通常表现为一方对另一方具有权力性的人际关系，例如：上下级关系、亲子关系、师生关系、医患关系等，无论年龄、职位，人人都有自己的权威，显然权威关系深刻的影响着每个人的生活。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">在长久的文化积淀下，很多时候动物被赋予了明显的性格特点和文化内涵，比如：人们常把身边勤恳工作的同事称为“勤劳的蜜蜂”。本投射测验运用多种“动物”形象及其关系状态投射出被试的权威关系模式。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">同时，每个人都具有多个性格和行为侧面，如：在工作时严肃认真，面对家人温柔幽默。因此，本测验选取14种常见的典型动物意象作为基本类型，将被试权威关系中的核心特质进行归类，符合测验本身“类型化”特点，在测试中要求被试选择与自身情况最接近的选项。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">权威关系心理模式通常是每个人通过幼年时与父母互动形成的，因此，父母是我们人生中的第一权威代表，而我们也通常在未来的生活中沿用与父母的人际模式来应对其他权威关系。如：父母总是非常严厉、命令式的管教，这样的子女长大后通常在上下级关系中也表现的顺从、沉默。</p><p style="font-size:16px;line-height:2em;text-indent:2em;">良好的人际关系有助于我们更好的工作、生活，获得更充分的自我发展，反之，则可能陷入举步维艰的困境。</p>`
	//汇总企业员工测试量表答题信息
	var (
		animalsFirst  []model.Animal
		animalsSecond []model.Animal
		attitudeThree []model.NormData
		attitudeFour  []model.NormData
		attitudeFive  []model.NormData
	)
	type AnswersCount struct {
		SubjectID       int
		SubjectSort     int
		SubjectAnswerID int
		AnswerSort      int
		OptionName      string
		Count           int
	}
	var answercounts []AnswersCount
	if err := db.Debug().Table("xy_staff_answer a").
		Joins("left join xy_staff_auswer_option b on a.id = b.staff_answer_id").
		Joins("left join xy_subject_answer c on b.subject_answer_id = c.id").
		Joins("JOIN xy_subject d ON a.gauge_id = d.gauge_id AND d.id = b.subject_id").
		Select("b.subject_id,d.sort as subject_sort,b.subject_answer_id,c.sort as answer_sort,c.option_name,COUNT(b.subject_answer_id)").
		Where("a.gauge_id = ? AND a.company_id = ? AND a.company_times = ?", gauge.ID, comID, comTimes).
		Group("b.subject_id ,d.sort, b.subject_answer_id,c.sort,c.option_name").
		Scan(&answercounts).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_staff_answer error!", file, line)
		return nil, err
	}

	for _, answerinfo := range answercounts {
		var animalTmp model.Animal
		var attitudeTmp model.NormData

		if answerinfo.SubjectSort == 1 {
			animalTmp.Name = answerinfo.OptionName
			animalTmp.Number = answerinfo.Count
			animalsFirst = append(animalsFirst, animalTmp)
		} else if answerinfo.SubjectSort == 2 {
			animalTmp.Name = answerinfo.OptionName
			animalTmp.Number = answerinfo.Count
			animalsSecond = append(animalsSecond, animalTmp)
		} else if answerinfo.SubjectSort == 3 {
			attitudeTmp.Name = answerinfo.OptionName
			attitudeTmp.Status = getNormStatus(answerinfo.SubjectID, answerinfo.SubjectAnswerID)
			attitudeThree = append(attitudeThree, attitudeTmp)
		} else if answerinfo.SubjectSort == 4 {
			attitudeTmp.Name = answerinfo.OptionName
			attitudeTmp.Status = getNormStatus(answerinfo.SubjectID, answerinfo.SubjectAnswerID)
			attitudeFour = append(attitudeFour, attitudeTmp)
		} else if answerinfo.SubjectSort == 5 {
			attitudeTmp.Name = answerinfo.OptionName
			attitudeTmp.Status = getNormStatus(answerinfo.SubjectID, answerinfo.SubjectAnswerID)
			attitudeFive = append(attitudeFive, attitudeTmp)
		}
	}

	return nil, nil
}

func getNormStatus(int, int) int {
	return 0
}

/*ConnectDataBase  connect database function*/
func ConnectDataBase() *gorm.DB {
	/*连接数据库*/
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/xyxjdata?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	/*关闭连接数据库*/
	//defer db.Close()
	return db
}
