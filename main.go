package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"heartcloud/config"
	"heartcloud/controller"
	"heartcloud/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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
	//添加中间件,验证用户登录信息
	router.Use(Middleware())
	router.POST("/staffreport", controller.GenerateStaffReport)
	router.POST("/companyreport", controller.GenerateCompanyReport)

	http.ListenAndServe(":8000", router)
}

/*Middleware 中间件 */
// func Middleware(c *gin.Context) {
// 	fmt.Println("this is a middleware!")
// }
/*Middleware 中间件 */
func Middleware() gin.HandlerFunc {
	fmt.Println("this is a middleware!")
	return func(c *gin.Context) {
		c.Keys, _ = verifyToken(c)
		fmt.Println("you don't known", c.Keys)
	}
}

//用户信息验证Token
func verifyToken(c *gin.Context) (map[string]interface{}, error) {
	/*读取redis获取token,通过token匹配用户信息BEGIN*/
	//获取Token
	authorization := c.GetHeader("Authorization")
	fmt.Printf("Headers is %v\n", c.Request.Header)
	// authorization := c.Request.Header.Get("authorization")
	fmt.Printf("@@@@@@@  1 23 authorization is :%s\n", authorization)
	tokenKey := model.AccessTokenPrefix + authorization
	//连接Redis
	conRedis, err := redis.Dial("tcp", config.RedisHost+config.RedisPort)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil, err
	}
	defer conRedis.Close()
	keyInfo, err := redis.Bytes(conRedis.Do("Get", tokenKey)) //获取客户端缓存信息
	if err != nil {
		log.Printf("登录信息已过期，请重新登陆！\n")
		return nil, err
	}
	fmt.Printf("@@@@@@   客户端信息:%v\n", keyInfo)

	var clientInfo map[string]interface{}
	err = json.Unmarshal(keyInfo, &clientInfo) //解析JSON字符串信息
	if err != nil {
		fmt.Printf("Unmarshal JSON error : %s\n", err)
		return nil, err
	}
	fmt.Println("####getkey ::", clientInfo)

	//获取缓存token
	cacheToken := clientInfo["access_token"]
	switch token := cacheToken.(type) {
	case string:
		if strings.Compare(token, authorization) != 0 {
			log.Printf("登录信息验证错误,请重新登录！\n")
			return nil, errors.New("登录信息验证错误,请重新登录！")
		}
	}

	//获取token过期时间(过期后无法得到token)
	// tokenExpireTime := clientInfo["expires_time"]
	/* isStaffInfo := clientInfo["client"]
	switch i := isStaffInfo.(type) { //switch type assertion
	case interface{}:
		fmt.Println("hahahahhaha", i)
	case map[string]interface{}:
		fmt.Println("map[string]interface{}", i)
	}
	fmt.Println(isStaffInfo) */
	// 获取员工登陆信息
	staffInfo, ok := clientInfo["client"].(map[string]interface{})
	if ok { //type assertion
		fmt.Println(staffInfo["staff_id"])
		fmt.Println(staffInfo["name"])
		fmt.Println(staffInfo["phone"])
		fmt.Println(staffInfo["company_id"])
		fmt.Println(staffInfo["company_name"])
	}

	/*读取redis获取token,通过token匹配用户信息END*/
	return staffInfo, nil
}
