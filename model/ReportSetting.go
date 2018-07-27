package model

import "github.com/jinzhu/gorm"

/*Reportsetting is the struct of table xy_report_setting */
type Reportsetting struct {
	gorm.Model
	GaugeID             int     `gorm:"type:int(11);DEFAULT NULL"`
	HideReportIntroduce int     `gorm:"type:smallint(1);Default '1'"`
	ReportIntroduce     string  `gorm:"type:text"`
	HideShowMethod      int     `gorm:"type:smallint(1);Default '1'"`
	ShowMethod          int     `gorm:"type:smallint(2)"`
	HideDescribe        int     `gorm:"type:smallint(1);Default '1'"`
	Describe            string  `gorm:"type:text"`
	HideComment         int     `gorm:"type:smallint(1);Default '1'"`
	Comment             string  `gorm:"type:text"`
	HideDimSuggest      int     `gorm:"type:smallint(1);Default '1'"`
	HideCliches         int     `gorm:"type:smallint(1);Default '1'"`
	Cliches             string  `gorm:"type:text"`
	ReferScore          float64 `gorm:"type:flaot(7,2);Default NULL"`
	ItScoreDesc         string  `gorm:"type:text"`
	GtScoreDesc         string  `gorm:"type:text"`
	/* CreatedTime         string
	UpdatedTime         string
	DeletedTime         string */
}
