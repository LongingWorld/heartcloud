package model

import "github.com/jinzhu/gorm"

/*ReportStaff is the struct of table xy_report_staff */
type ReportStaff struct {
	gorm.Model
	StaffAnswerID       int     `gorm:"type:int(11);Default NULL"`
	Name                string  `gorm:"type:varchar(255);Default NULL"`
	HideReportIntroduce int     `gorm:"type:smallint(1);Default '1'"`
	Introduce           string  `gorm:"type:text"`
	HideDescribe        int     `gorm:"type:smallint(1);Default '1'"`
	Describe            string  `gorm:"type:text"`
	HideShowMethod      int     `gorm:"type:smallint(1);Default '1'"`
	ShowMethod          int     `gorm:"type:smallint(2)"`
	HideCliches         int     `gorm:"type:smallint(1);Default '1'"`
	Cliches             string  `gorm:"type:text"`
	HideComment         int     `gorm:"type:smallint(1);Default '1'"`
	Comment             string  `gorm:"type:text"`
	HideDimSuggest      int     `gorm:"type:smallint(1);Default '1'"`
	GaugeID             int     `gorm:"type:int(11);DEFAULT NULL"`
	StaffID             int     `gorm:"type:int(11);DEFAULT NULL"`
	StaffName           string  `gorm:"type:varchar(255);DEFAULT NULL"`
	StaffAge            int     `gorm:"type:int(11);DEFAULT NULL"`
	Position            string  `gorm:"type:varchar(255);DEFAULT NULL"`
	Marriage            int     `gorm:"type:smallint(6);DEFAULT NULL"`
	CompanyID           int     `gorm:"type:int(11);DEFAULT NULL"`
	CompanyName         string  `gorm:"type:varchar(255);DEFAULT NULL"`
	Number              string  `gorm:"type:char(12);DEFAULT NULL"`
	Status              int     `gorm:"type:tinyint(1);DEFAULT '1'"`
	TotalScore          float32 `gorm:"type:float(8,2);DEFAULT NULL"`
	TemplateID          int     `gorm:"type:smallint(2);DEFAULT NULL"`
	GenerateDate        string  `gorm:"type:timestamp;NULL DEFAULT NULL"`
}
