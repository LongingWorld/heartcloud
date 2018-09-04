package controller

import (
	"encoding/json"
	"fmt"
	"heartcloud/database"
	"heartcloud/model"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*AnswersCount is the struct of Statistics of answers*/
type AnswersCount struct {
	SubjectID       int
	SubjectSort     int
	SubjectAnswerID int
	AnswerSort      int
	OptionName      string
	Count           int
}

/*GenerateCompanyReport function generate company reports*/
func GenerateCompanyReport(c *gin.Context) {
	fmt.Println("@@@@@@@GenerateCompanyReport()Begin@@@@@@@")
	//验证登录Token信息，并获取用户信息
	/* companyInfo, err := verifyToken(c)
	if err != nil {
		log.Printf("验证Token信息失败！\n")
		c.JSON(500, "Token失效")
		return
	}
	fmt.Println(companyInfo) */
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
		fmt.Printf("?????value is %s?????\n", value)
		gaugeID, _ := strconv.Atoi(value)
		gaugeIDs = append(gaugeIDs, gaugeID)
		fmt.Printf("@@@@@@  value is %s,gaugeID is %d \n", value, gaugeID)
	}
	fmt.Printf("@@@@@@   gaugeIDs is %v\n", gaugeIDs)
	//连接数据库
	db := database.ConnectDataBase()
	/*关闭连接数据库*/
	defer db.Close()

	//获取量表列表信息
	var gauges []model.Gauge
	if err := db.Debug().Table("xy_gauge").
		Where("id in (?)", gaugeIDs).
		Scan(&gauges).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_gauge error!", file, line)
		c.JSON(500, "系统异常")
		return
	}

	/* 获取report_company_id插入的企业报告ID */
	/*获取企业报告ID*/
	repCompID := c.PostForm("report_company_id")
	reportCompanyID, _ := strconv.Atoi(repCompID)
	fmt.Printf("@@@@@@   reportCompanyID is : %d\n", reportCompanyID)
	//企业报告report_company_data表中的report_data以及 report_data_api字段
	type ReportData struct {
		ReportData    string
		ReportDataAPI string
	}
	var (
		reportDataStr ReportData
		repData       map[string]interface{}
		repDataAPI    map[string]interface{}
	)

	if err := db.Debug().Table("xy_report_company_data").
		Where("report_company_id = ?", reportCompanyID).
		Scan(&reportDataStr).Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Select Table xy_gauge error!", file, line)
		return
	}

	fmt.Printf("@@@@@@Original   reportdata is: \n%v\n", reportDataStr.ReportData)
	//解析report_company_data.report_data to map
	if err := json.Unmarshal([]byte(reportDataStr.ReportData), &repData); err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:解析JSON字符串出错", file, line)

		return
	}
	//fmt.Printf("@@@@@@Original   reportdata is: \n%v\n", repData)
	//解析report_company_data.report_data_api to map
	if err := json.Unmarshal([]byte(reportDataStr.ReportDataAPI), &repDataAPI); err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:解析JSON字符串出错", file, line)
		return
	}
	//fmt.Printf("@@@@@@Original   reportdataapi is: \n%v\n", repDataAPI)
	/*开启事务*/
	tx := db.Begin()

	for _, gaugeinfo := range gauges {
		if gaugeinfo.TemplateID == 1 || gaugeinfo.TemplateID == 2 || gaugeinfo.TemplateID == 3 {
			continue
		} else if gaugeinfo.TemplateID == 4 {
			repComData, err := createAuthorityRelationComReportData(tx, gaugeinfo, companyID, distributeTime)
			if err != nil {
				c.JSON(500, "系统异常")
			}

			repData["template4"] = repComData
			repDataAPI["template4"] = repComData
			//fmt.Printf("######  template4 data is :\n %v\n", repData)
		} else if gaugeinfo.TemplateID == 5 {
			repComData, err := createChronicFatiguesComReportData(tx, gaugeinfo, companyID, distributeTime)
			if err != nil {
				c.JSON(500, "系统异常")
			}

			repData["template5"] = repComData
			repDataAPI["template5"] = repComData
			//fmt.Printf("######  template4 data is :\n %v\n", repData)
		} else if gaugeinfo.TemplateID == 6 {
			repComData, err := createEgoStateCompanyReportData(tx, int(gaugeinfo.ID), companyID, distributeTime)
			if err != nil {
				c.JSON(500, "系统异常")
			}

			repData["template6"] = repComData
			repDataAPI["template6"] = repComData
			//fmt.Printf("######  template4 data is :\n %v\n", repData)
		} else if gaugeinfo.TemplateID == 7 {
			repComData, err := createDSQComReportData(tx, int(gaugeinfo.ID), companyID, distributeTime)
			if err != nil {
				c.JSON(500, "系统异常")
			}

			repData["template7"] = repComData
			repDataAPI["template7"] = repComData
			//fmt.Printf("######  template4 data is :\n %v\n", repData)
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
		c.JSON(500, "系统异常")
		return
	}
	// fmt.Printf("@@@@@@   reportCompanyData is:\n %s\n", string(reportCompanyData))
	fmt.Printf("@@@@@@   reportCompanyData byte  is:\n %d\n", len(reportCompanyData))

	//map to json
	reportCompanyDataAPI, err := json.Marshal(&repDataAPI)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:解析JSON字符串出错", file, line)
		c.JSON(500, "系统异常")
		return
	}
	//fmt.Printf("@@@@@@   reportCompanyDataAPI is:\n %s\n", string(reportCompanyDataAPI))

	//更新xy_report_company_data.report_data数据
	if err := tx.Debug().Table("xy_report_company_data").
		Where("report_company_id = ?", reportCompanyID).
		Updates(map[string]interface{}{"report_data": string(reportCompanyData), "report_data_api": string(reportCompanyDataAPI)}).
		Error; err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Printf("%s:%d:Update Table xy_report_company_data error!", file, line)
		tx.Rollback()
		c.JSON(500, "系统异常")
		return
	}
	//提交事物
	tx.Commit()
	c.JSON(http.StatusOK, "success")
	fmt.Println("@@@@@@@GenerateCompanyReport()end@@@@@@@")
	return
}
