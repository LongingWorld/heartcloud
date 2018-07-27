package model

import "time"

/*Model is the Basic model definition */
type Model struct {
	ID        int        `gorm:"type:int(11) unsigned ;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"type:timestamp;CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"type:timestamp;DEFAULT NULL"`
	DeletedAt *time.Time `gorm:"type:timestamp;DEFAULT NULL"`
}
