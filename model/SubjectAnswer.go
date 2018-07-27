package model

import "time"

/*SubjectAnswer is the struct of table xy_subject_answer */
type SubjectAnswer struct {
	ID           int       `gorm:"type:int(11) unsigned ;primary_key;AUTO_INCREMENT"`
	SubjectID    int       `gorm:"type:int(11);Default NULL"`
	OptionName   string    `gorm:"type:text"`
	Image        string    `gorm:"type:varchar(255);Default NULL"`
	Fraction     int       `gorm:"type:int(3);Default NULL"`
	MappingValue string    `gorm:"type:varchar(255);Default NULL"`
	Sort         int       `gorm:"type:smallint(2);Default NULL"`
	CreatedAt    time.Time `gorm:"type:timestamp;CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:timestamp;DEFAULT NULL"`
}
