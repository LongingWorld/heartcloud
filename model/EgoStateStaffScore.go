package model

//EgoStateStaffScoreTable  is the struct of the table xy_egostate_staff_dim_score
type EgoStateStaffScoreTable struct {
	EgoStateAnsID         int `gorm:"type:int(18);primary_key"`
	StaffID               int `gorm:"type:int(11)"`
	CompanyID             int `gorm:"type:int(11)"`
	CompanyTimes          int `gorm:"type:int(11)"`
	GaugeID               int `gorm:"type:int(11)"`
	PosControlParentScore int `gorm:"type:int(3)"`
	NegControlParentScore int `gorm:"type:int(3)"`
	PosCareParentScore    int `gorm:"type:int(3)"`
	NegCareParentScore    int `gorm:"type:int(3)"`
	AdultScore            int `gorm:"type:int(3)"`
	PosFreeChildScore     int `gorm:"type:int(3)"`
	NegFreeChildScore     int `gorm:"type:int(3)"`
	PosObeyChildScore     int `gorm:"type:int(3)"`
	NegObeyChildScore     int `gorm:"type:int(3)"`
	RebelChildScore       int `gorm:"type:int(3)"`
}
