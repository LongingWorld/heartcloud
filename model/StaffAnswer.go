package model

import "github.com/jinzhu/gorm"

/*StaffAnswer is the struct of table xy_staff_answer */
type StaffAnswer struct {
	gorm.Model
	ServiceUseStaffID int    `gorm:"type:int(11);Default NULL"`
	CompanyID         int    `gorm:"type:int(11);Default NULL"`
	GaugeID           int    `gorm:"type:int(11);Default NULL"`
	StartTime         string `gorm:"type:timestamp;Default NULL"`
	EndTime           string `gorm:"type:timestamp;Default NULL"`
	Score             int    `gorm:"type:smallint(4);Default NULL"`
	StaffID           int    `gorm:"type:int(11);Default NULL"`
	IsFinish          int    `gorm:"type:smallint(1);Default '0'"`
	CompanyTimes      int    `gorm:"type:int(11);Default '1'"`
}
