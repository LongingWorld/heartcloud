package main

import (
	"fmt"
	"heartcloud/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* type postRequst struct {
	accessToken     string
	gaugeID         string
	subjectsAnswers string
}
*/

func main() {
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()
	//添加中间件
	router.Use(Middleware)
	router.POST("/staffreport", controller.GenerateStaffReport)
	router.POST("/companyreport", controller.GenerateCompanyReport)

	http.ListenAndServe(":8000", router)
}

/*Middleware 中间件 */
func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}
