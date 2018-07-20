package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const AccessTokenPrefix = "accessToken_"

// 基本模型的定义
type Model struct {
	ID        int        `gorm:"type:int(11) unsigned ;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"type:timestamp;CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"type:timestamp;DEFAULT NULL"`
	DeletedAt *time.Time `gorm:"type:timestamp;DEFAULT NULL"`
}

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

type Reportsetting struct {
	gorm.Model
	GaugeID             int     `gorm:"type:int(11);DEFAULT NULL"`
	HideReportIntroduce int     `gorm:"type:smallint(1);Default '1'"`
	ReportIntroduce     string  `gorm:"type:text"`
	HideShowMethod      int     `gorm:"type:smallint(1);Default '1'"`
	ShowMethod          int     `gorm:"type:smallint(2)"`
	HideDescribe        int     `gorm:"type:smallint(1);Default '1'"`
	Describe            string  `gorm:"type:text"`
	HideComment         int     `gorm:"type:smallint(1);Default '1'"`
	Comment             string  `gorm:"type:text"`
	HideDimSuggest      int     `gorm:"type:smallint(1);Default '1'"`
	HideCliches         int     `gorm:"type:smallint(1);Default '1'"`
	Cliches             string  `gorm:"type:text"`
	ReferScore          float64 `gorm:"type:flaot(7,2);Default NULL"`
	ItScoreDesc         string  `gorm:"type:text"`
	GtScoreDesc         string  `gorm:"type:text"`
	/* CreatedTime         string
	UpdatedTime         string
	DeletedTime         string */
}

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

type StaffAnswer struct {
	gorm.Model
	ServiceUseStaffID int    `gorm:"type:int(11);Default NULL"`
	CompanyID         int    `gorm:"type:int(11);Default NULL"`
	GaugeID           int    `gorm:"type:int(11);Default NULL"`
	StartTime         string `gorm:"type:timestamp;Default NULL"`
	EndTime           string `gorm:"type:timestamp;Default NULL"`
	Score             int    `gorm:"type:smallint(4);Default NULL"`
	StaffID           int    `gorm:"type:int(11);Default NULL"`
	IsFinish          int    `gorm:"type:smallint(1);Default '0'"`
	CompanyTimes      int    `gorm:"type:int(11);Default '1'"`
	/* CreatedTime       string
	UpdatedTime       string
	DeletedTime       string */
}

type Subject struct {
	gorm.Model
	GaugeID     int    `gorm:"type:int(11);Default NULL"`
	SubjectName string `gorm:"type:text"`
	Sort        int    `gorm:"type:int(11);Default NULL"`
	Number      int    `gorm:"type:int(3);Default NULL"`
	/* CreatedTime string
	UpdatedTime string
	DeletedTime string */
}

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

type ReportStaff struct {
	gorm.Model
	StaffAnswerID       int    `gorm:"type:int(11);Default NULL"`
	Name                string `gorm:"type:varchar(255);Default NULL"`
	HideReportIntroduce int    `gorm:"type:smallint(1);Default '1'"`
	Introduce           string `gorm:"type:text"`
	HideDescribe        int    `gorm:"type:smallint(1);Default '1'"`
	// ItScoreDesc         string  `gorm:"type:text"`
	Describe       string  `gorm:"type:text"`
	HideShowMethod int     `gorm:"type:smallint(1);Default '1'"`
	ShowMethod     int     `gorm:"type:smallint(2)"`
	HideCliches    int     `gorm:"type:smallint(1);Default '1'"`
	Cliches        string  `gorm:"type:text"`
	HideComment    int     `gorm:"type:smallint(1);Default '1'"`
	Comment        string  `gorm:"type:text"`
	HideDimSuggest int     `gorm:"type:smallint(1);Default '1'"`
	GaugeID        int     `gorm:"type:int(11);DEFAULT NULL"`
	StaffID        int     `gorm:"type:int(11);DEFAULT NULL"`
	StaffName      string  `gorm:"type:varchar(255);DEFAULT NULL"`
	StaffAge       int     `gorm:"type:int(11);DEFAULT NULL"`
	Position       string  `gorm:"type:varchar(255);DEFAULT NULL"`
	Marriage       int     `gorm:"type:smallint(6);DEFAULT NULL"`
	CompanyID      int     `gorm:"type:int(11);DEFAULT NULL"`
	CompanyName    string  `gorm:"type:varchar(255);DEFAULT NULL"`
	Number         string  `gorm:"type:char(12);DEFAULT NULL"`
	Status         int     `gorm:"type:tinyint(1);DEFAULT '1'"`
	TotalScore     float32 `gorm:"type:float(8,2);DEFAULT NULL"`
	TemplateID     int     `gorm:"type:smallint(2);DEFAULT NULL"`
	GenerateDate   string  `gorm:"type:timestamp;NULL DEFAULT NULL"`
	/* 	CreatedTime         string
	   	UpdatedTime         string
	   	DeletedTime         string */
}

type ReportStaffData struct {
	ReportStaffID   int    `gorm:"type:int(11);DEFAULT NULL"`
	ReportData      string `gorm:"type:mediumtext"`
	ReportDataExtra string `gorm:"type:mediumtext"`
}

type NormDetail struct {
	ID             int    `json:"norm_explain_id"`
	Name           string `json:"name"`
	ScoreIntroduce string `json:"score_introduce"`
	CoachProposal  string `json:"coach_proposal"`
	//IsCheck        int    `json:"is_check"`
}
type ExplainsDetail struct {
	MapID       int    `json:"norm_explain_id"`
	OptionName  string `json:"name"`
	MapExplain  string `json:"score_introduce"`
	MapProposal string `json:"coach_proposal"`
}
type StaffDimension struct {
	StaffID        int              `json:"staff_id"`
	DimensionID    int              `json:"dimension_id"`
	Score          int              `json:"score"`
	Extra          string           `json:"extra"`
	DimName        string           `json:"dim_name"`
	DimDesc        string           `json:"dim_desc"`
	DimSuggest     string           `json:"dim_suggest"`
	Sort           int              `json:"sort"`
	Formula        string           `json:"formula"`
	NormID         int              `json:"norm_id"`
	NormName       string           `json:"norm_name"`
	NormType       int              `json:"norm_type"`
	ReferenctValue string           `json:"reference_value"`
	ExplainsDetail []ExplainsDetail `json:"explains_detail"`
}

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
