package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	fmt.Println("@@@@@@@GenerateCompanyReport()end@@@@@@@")
}
