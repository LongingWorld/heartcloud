package model

//EgoStateInfoTable  is the struct of the table xy_ego_state_info
type EgoStateInfoTable struct {
	EgoID          string `gorm:"type:varchar(8);primary_key"`
	EgoName        string `gorm:"type:varchar(8);primary_key"`
	EgoMin         int    `gorm:"type:tinyint(2);primary_key"`
	EgoMax         int    `gorm:"type:varchar(2);primary_key"`
	EgoSqe         int    `gorm:"type:tinyint(1);primary_key"`
	EgoBriefName   string `gorm:"type:varchar(45);Default NULL"`
	EgoBriefInfo   string `gorm:"type:text"`
	EgoResultBegin string `gorm:"type:text"`
	EgoResultEnd   string `gorm:"type:text"`
	EgoAlwaysTitle string `gorm:"type:text"`
	EgoAlwaysDesc  string `gorm:"type:text"`
	EgoRarelyTitle string `gorm:"type:text"`
	EgoRarelyDesc  string `gorm:"type:text"`
	Remark1        string `gorm:"type:text"`
	Remark2        string `gorm:"type:text"`
}
