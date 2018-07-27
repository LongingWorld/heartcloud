package model

import "github.com/jinzhu/gorm"

/*Subject is the struct of table xy_subject */
type Subject struct {
	gorm.Model
	GaugeID     int    `gorm:"type:int(11);Default NULL"`
	SubjectName string `gorm:"type:text"`
	Sort        int    `gorm:"type:int(11);Default NULL"`
	Number      int    `gorm:"type:int(3);Default NULL"`
}
