package model

import "github.com/jinzhu/gorm"

/*Company is the struct of table xy_company */
type Company struct {
	gorm.Model
	Name           string `gorm:"type:varchar(255);Default NULL"`
	UserName       string `gorm:"type:varchar(255);Default NULL"`
	Password       string `gorm:"type:char(32);Default NULL"`
	Phone          string `gorm:"type:char(11);Default NULL"`
	ViewID         int    `gorm:"type:int(255);Default NULL"`
	ContractNumber string `gorm:"type:varchar(255);Default NULL"`
	Remarks        string `gorm:"type:varchar(255);Default NULL"`
	StartTime      string `gorm:"type:timestamp;Default NULL"`
	EndTime        string `gorm:"type:timestamp;Default NULL"`
	AdminID        int    `gorm:"type:int(11);Default NULL"`
	/* CreatedTime    string
	UpdatedTime    string
	DeletedTime    string */
}
