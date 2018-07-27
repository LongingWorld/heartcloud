package model

/*AccessTokenPrefix is The prefix of Token*/
const AccessTokenPrefix = "accessToken_"

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

/*CompReportDetail is the response of report_company template_id=4*/
type CompReportDetail struct {
	GaugeID       int      `json:"id"`
	GaugeName     string   `json:"name"`
	GaugeShowName string   `json:"show_name"`
	TemplateID    int      `json:"template_id"`
	Introduction  string   `json:"introduction"`
	Section1      Section1 `json:"Section1"`
	Section2      Section2 `json:"Section2"`
	Section3      Section3 `json:"Section3"`
	Section4      Section4 `json:"Section4"`
	Section5      Section5 `json:"Section5"`
	Section6      Section6 `json:"Section6"`
}

/*Section1 is the datastruct  of the first question */
type Section1 struct {
	Data      []Animal `json:"data"`
	DescAnaly DescAnalysis
}

/*Section2 is the datestruct of the second question */
type Section2 struct {
	Data      []Animal `json:"data"`
	DescAnaly DescAnalysis
}

/*Section3 is the datastruct of the third question*/
type Section3 struct {
	Data      []NormData `json:"data"`
	DescAnaly DescAnalysis
}

/*Section4 is the datastruct of the fourth question*/
type Section4 struct {
	Data      []NormData `json:"data"`
	DescAnaly DescAnalysis
}

/*Section5 is the datastruct of the fifth question*/
type Section5 struct {
	Data      []NormData `json:"data"`
	DescAnaly DescAnalysis
}

/*Section6 is the datastruct of the General health instructions*/
type Section6 struct {
	Data      []NormData `json:"data"`
	DescAnaly DescAnalysis
}

/*DescAnalysis is the struct of describe and analysis of those questions */
type DescAnalysis struct {
	Describe string `json:"desc"`
	Analysis string `json:"analysis"`
}

/*Animals is the struct of the first and second question answers*/
/* type Animals struct {
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
} */

/*Animal is the comment of the first and second question*/
type Animal struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

// /*SectionTwo is the struct of the third question*/
// type SectionTwo struct {
// 	Data      AttitudeToSuperior `json:"data"`
// 	DescAnaly DescAnalysis
// }
// type AttitudeToSuperior struct {
// 	Adore     NormData `json:"adore"`
// 	Attitueds AttitudeMembers
// }
// type SectionThree struct {
// 	Data      AttitudeToStaff `json:"data"`
// 	DescAnaly DescAnalysis
// }
// type AttitudeToStaff struct {
// 	Appreciate NormData `json:"appreciate"`
// 	Attitueds  AttitudeMembers
// }
// type AttitudeMembers struct {
// 	Accept  NormData `json:"accept"`
// 	Notcare NormData `json:"notcare"`
// 	Deny    NormData `json:"deny"`
// 	Hate    NormData `json:"hate"`
// 	Fear    NormData `json:"fear"`
// }

/*NormData is the normal data struct */
type NormData struct {
	Name    string  `json:"name"`
	Status  int     `json:"status"`
	Persent float64 `json:"persent"`
}

/* type SectionFour struct {
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
} */
