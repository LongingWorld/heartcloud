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

/*StaffAnswer is the struct of table xy_staff_answer */
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

/*ReportStaffData is mapping  xy_report_staff_data */
type ReportStaffData struct {
	ReportStaffID   int    `gorm:"type:int(11);DEFAULT NULL"`
	ReportData      string `gorm:"type:mediumtext"`
	ReportDataExtra string `gorm:"type:mediumtext"`
}

/*NormDetail is the norm explain information */
type NormDetail struct {
	ID             int    `json:"norm_explain_id"`
	Name           string `json:"name"`
	ScoreIntroduce string `json:"score_introduce"`
	CoachProposal  string `json:"coach_proposal"`
	//IsCheck        int    `json:"is_check"`
}

/*ExplainsDetail is the norm explain information  */
type ExplainsDetail struct {
	MapID       int    `json:"norm_explain_id"`
	OptionName  string `json:"name"`
	MapExplain  string `json:"score_introduce"`
	MapProposal string `json:"coach_proposal"`
}

/*StaffDimension is the response of report_staff */
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

/*ReportCompany is the struct of table xy_report_company */
type ReportCompany struct {
	gorm.Model
	ReportName   string `gorm:"type:varchar(128);Default NULL"`
	BgDesc       string `gorm:"type:text"`
	Introduce    string `gorm:"type:text"`
	Number       string `gorm:"type:varchar(12);DEFAULT NULL "`
	CompanyID    int    `gorm:"type:int(11);DEFAULT NULL"`
	GaugeIds     string `gorm:"type:text"`
	Status       int    `gorm:"type:tinyint(1);DEFAULT NULL"`
	Propocal     string `gorm:"type:text"`
	GenerateDate string `gorm:"type:timestamp;NULL DEFAULT NULL"`
}

/*ReportCompanyData is mapping  xy_report_company_data */
type ReportCompanyData struct {
	ReportCompanyID int    `gorm:"type:int(11);DEFAULT NULL"`
	ReportData      string `gorm:"type:mediumtext"`
	URL             string `gorm:"type:varchar(255);DEFAULT NULL"`
	ReportDataAPI   string `gorm:"type:mediumtext"`
}

/*CompReportDetail is the response of report_company template_id=4*/
type CompReportDetail struct {
	GaugeID       int          `json:"id"`
	GaugeName     string       `json:"name"`
	GaugeShowName string       `json:"show_name"`
	TemplateID    int          `json:"template_id"`
	Introduction  string       `json:"section1"`
	Section2      SectionOne   `json:"Section2"`
	Section3      SectionOne   `json:"section3"`
	Section4      SectionTwo   `json:"Section4"`
	Section5      SectionThree `json:"section5"`
	Section6      SectionFour  `json:"Section6"`
	Section7      SectionFive  `json:"section7"`
}

/*SectionOne is the introduction of the company report */
type SectionOne struct {
	Data      Animals `json:"data"`
	DescAnaly DescAnalysis
}

/*DescAnalysis is the struct of describe and analysis of those questions */
type DescAnalysis struct {
	Describe string `json:"desc"`
	Analysis string `json:"analysis"`
}

/*Animals is the struct of the first and second question answers*/
type Animals struct {
	Tiger  Animal `json:"tiger"`
	Dog    Animal `json:"dog"`
	Ant    Animal `json:"ant"`
	Wolf   Animal `json:"wolf"`
	Fox    Animal `json:"fox"`
	Spider Animal `json:"spider"`
	Lion   Animal `json:"lion"`
	Cow    Animal `json:"cow"`
	Sheep  Animal `json:"sheep"`
	Horse  Animal `json:"horse"`
	Rabbit Animal `json:"rabbit"`
	Bee    Animal `json:"bee"`
	Cayman Animal `json:"cayman"`
	Mouse  Animal `json:"mouse"`
}

/*Animal is the comment of the first and second question*/
type Animal struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

/*SectionTwo is the struct of the third question*/
type SectionTwo struct {
	Data      AttitudeToSuperior `json:"data"`
	DescAnaly DescAnalysis
}
type AttitudeToSuperior struct {
	Adore     NormData `json:"adore"`
	Attitueds AttitudeMembers
}
type SectionThree struct {
	Data      AttitudeToStaff `json:"data"`
	DescAnaly DescAnalysis
}
type AttitudeToStaff struct {
	Appreciate NormData `json:"appreciate"`
	Attitueds  AttitudeMembers
}
type AttitudeMembers struct {
	Accept  NormData `json:"accept"`
	Notcare NormData `json:"notcare"`
	Deny    NormData `json:"deny"`
	Hate    NormData `json:"hate"`
	Fear    NormData `json:"fear"`
}

type NormData struct {
	Name    string  `json:"name"`
	Status  int     `json:"status"`
	Persent float64 `json:"persent"`
}

type SectionFour struct {
	Data      CooperationIndex `json:"data"`
	DescAnaly DescAnalysis
}
type CooperationIndex struct {
	FirstQuadrant  NormData `json:"firstQuadrant"`
	SecondQuadrant NormData `json:"secondQuadrant"`
	ThirdQuadrant  NormData `json:"thirdQuadrant"`
	FourthQuadrant NormData `json:"fourthQuadrant"`
}

type SectionFive struct {
	Data      RelationshipStatus `json:"data"`
	DescAnaly DescAnalysis
}
type RelationshipStatus struct {
	Level1 NormData `json:"level_1"`
	Level2 NormData `json:"level_2"`
	Level3 NormData `json:"level_3"`
	Level4 NormData `json:"level_4"`
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
