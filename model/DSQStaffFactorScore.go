package model

//DSQStaffFactorScore 防御方式-员工防御因子得分信息表
type DSQStaffFactorScore struct {
	DsqID                      int `gorm:"type:int(18);primary_key"`
	StaffID                    int `gorm:"type:int(11)"`
	CompanyID                  int `gorm:"type:int(11)"`
	CompanyTimes               int `gorm:"type:int(11)"`
	GaugeID                    int `gorm:"type:int(11)"`
	ConsealmentScore           int `gorm:"type:int(2)"`
	MatureScore                int `gorm:"type:int(2)"`
	IntermediateScore          int `gorm:"type:int(2)"`
	NotMatureScore             int `gorm:"type:int(2)"`
	SublimeScore               int `gorm:"type:int(2)"`
	HumorScore                 int `gorm:"type:int(2)"`
	ReactionformationScore     int `gorm:"type:int(2)"`
	RelieveScore               int `gorm:"type:int(2)"`
	DebarbScore                int `gorm:"type:int(2)"`
	RetionnaliseScore          int `gorm:"type:int(2)"`
	FalseAltruismScore         int `gorm:"type:int(2)"`
	HalfIncappableScore        int `gorm:"type:int(2)"`
	InsulateScore              int `gorm:"type:int(2)"`
	IdenticalTrendScore        int `gorm:"type:int(2)"`
	DenyScore                  int `gorm:"type:int(2)"`
	ConsumptionTendenciesScore int `gorm:"type:int(2)"`
	ExpectScore                int `gorm:"type:int(2)"`
	AttiliationScore           int `gorm:"type:int(2)"`
	CurbScore                  int `gorm:"type:int(2)"`
	DepressScore               int `gorm:"type:int(2)"`
	DartScore                  int `gorm:"type:int(2)"`
	PassiveAttackScore         int `gorm:"type:int(2)"`
	SubconsciousScore          int `gorm:"type:int(2)"`
	ComplainScore              int `gorm:"type:int(2)"`
	FantasyScore               int `gorm:"type:int(2)"`
	SplitScore                 int `gorm:"type:int(2)"`
	SomatizationScore          int `gorm:"type:int(2)"`
	FlinchScore                int `gorm:"type:int(2)"`
}
