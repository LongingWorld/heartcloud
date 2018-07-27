package model

import "time"

/*SystemConf is the struct of table xy_system_conf */
type SystemConf struct {
	ParaID      string     `gorm:"type:varchar(8);primary_key"`
	ParaName    string     `gorm:"type:varchar(255);NOT NULL"`
	ParaType    int        `gorm:"type:int(2);DEFAULT NULL"`
	ParaStatus  int        `gorm:"type:int(1);DEFAULT NULL"`
	ParaExplain string     `gorm:"type:varchar(255);DEFAULT NULL"`
	ParaConf    string     `gorm:"type:text"`
	Remark1     string     `gorm:"type:varchar(255) DEFAULT NULL"`
	Remark2     string     `gorm:"type:varchar(1024) DEFAULT NULL"`
	CreatedAt   time.Time  `gorm:"type:timestamp;CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;DEFAULT NULL"`
	DeletedAt   *time.Time `gorm:"type:timestamp;DEFAULT NULL"`
}
