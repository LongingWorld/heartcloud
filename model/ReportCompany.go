package model

import "github.com/jinzhu/gorm"

/*ReportCompany is the struct of table xy_report_company */
type ReportCompany struct {
	gorm.Model
	ReportName   string `gorm:"type:varchar(128);Default NULL"`
	BgDesc       string `gorm:"type:text"`
	Introduce    string `gorm:"type:text"`
	Number       string `gorm:"type:varchar(12);DEFAULT NULL "`
	CompanyID    int    `gorm:"type:int(11);DEFAULT NULL"`
	GaugeIds     string `gorm:"type:text"`
	Status       int    `gorm:"type:tinyint(1);DEFAULT NULL"`
	Propocal     string `gorm:"type:text"`
	GenerateDate string `gorm:"type:timestamp;NULL DEFAULT NULL"`
}
