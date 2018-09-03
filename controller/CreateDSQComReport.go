package controller

import (
	"fmt"
	"heartcloud/model"

	"github.com/jinzhu/gorm"
)

func createDSQComReportData(db *gorm.DB, gaugeID int, comID int, comTimes int) (model.DSQCompanyReportData, error) {
	fmt.Println("**********************createDSQComReportData BEGIN************************")
	fmt.Println("**********************createDSQComReportData END************************")
	return model.DSQCompanyReportData{}, nil
}
