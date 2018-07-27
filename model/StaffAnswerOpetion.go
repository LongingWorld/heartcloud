package model

import "github.com/jinzhu/gorm"

/*StaffAnswerOpetion is the struct of table xy_staff_auswer_option */
type StaffAnswerOpetion struct {
	gorm.Model
	StaffAnswerID   int `gorm:"type:int(11);Default NULL"`
	SubjectID       int `gorm:"type:int(11);Default NULL"`
	SubjectAnswerID int `gorm:"type:int(11);Default NULL"`
	StaffID         int `gorm:"type:int(11);Default NULL"`
	/* CreatedTime     string
	UpdatedTime     string
	DeletedTime     string */
}
