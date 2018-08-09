package database

import (
	"fmt"
	"heartcloud/config"

	"github.com/jinzhu/gorm"
)

/*ConnectDataBase : Use the function to connect database */
func ConnectDataBase() *gorm.DB {
	/*连接数据库*/
	fmt.Printf("###### ConnectDataBase %s%s", config.DbHost, config.DbPort)
	db, err := gorm.Open("mysql", "root:@tcp("+config.DbHost+config.DbPort+")/xyxjdata?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	/*关闭连接数据库*/
	//defer db.Close()
	return db
}
