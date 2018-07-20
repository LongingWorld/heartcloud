package handler_test

import (
	"log"
	"strconv"
	"testing"

	"xinyun/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestPostHandler(t *testing.T) {

	/* router := gin.Default()
	router.POST("/first", handler.PostHandler)

	http.ListenAndServe(":8000", router) */
	var c *gin.Context

	/*连接数据库*/
	db, err := gorm.Open("mysql", "root:zzj210@tcp(127.0.0.1:3308)/world?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	/*关闭连接数据库*/
	defer db.Close()

	var gauge model.Gauge
	// var company Company
	//var reportSet model.Reportsetting

	/*获取量表ID*/
	gaugeID := c.PostForm("gauge_id")
	//db.First(&gauge, "id = ?", gaugeID)
	gint, _ := strconv.Atoi(gaugeID)
	db.Where("id = ?", gint).Find(&gauge)
	var count int
	db.Model(&model.Gauge{}).Where("id = ?", gint).Count(&count) //获取记录数
	if count == 0 {
		log.Printf("量表%s不存在", gaugeID)
		return
		//fmt.Errorf("量表%s不存在", gaugeID)
	}
}
