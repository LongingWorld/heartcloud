package model

import "github.com/jinzhu/gorm"

/*Gauge is the struct 0of table xy_gauge */
type Gauge struct {
	gorm.Model
	Name           string `gorm:"type:varchar(100);DEFAULT NULL"`
	ShowName       string `gorm:"type:varchar(100;DEFAULT NULL)"`
	Describe       string `gorm:"type:text"`
	CategoryID     int    `gorm:"type:int(11);DEFAULT NULL"`
	IsRandom       int    `gorm:"type:smallint(1);DEFAULT NULL"`
	CompletionTime int    `gorm:"type:smallint(3);DEFAULT NULL"`
	Guidance       string `gorm:"type:text"`
	Status         int    `gorm:"type:smallint(1);Default '1'"`
	TemplateID     int    `gorm:"type:smallint(3);DEFAULT NULL"`
}
