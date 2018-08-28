package model

//ChronicFatigueComRepData is the struct of the ChronicFatigue gauger of comany report
type ChronicFatigueComRepData struct {
	GaugeID       int                     `json:"id"`
	GaugeName     string                  `json:"name"`
	GaugeShowName string                  `json:"show_name"`
	TemplateID    int                     `json:"template_id"`
	BriefInfo     ChronicFatigueBriefInfo `json:"brief_info"`
	DataAnalysis  ChronicFatigueData      `json:"data_analysis"`
}

//ChronicFatigueBriefInfo is the member of ChronicFatigueComRepData
type ChronicFatigueBriefInfo struct {
	BriefInfo string      `json:"brief_doc"`
	Classify  []ClassInfo `json:"classify"`
}

//ClassInfo is the memeber of ChronicFatigueBriefInfo
type ClassInfo struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

//ChronicFatigueData is the member of ChronicFatigueComRepData
type ChronicFatigueData struct {
	DegreeOfFatigue ChronicFatigugeRankStatus
	FatigueFactor   FatigueFactor
	SexFactor       SexCompare
	AgeFactor       AgeCompare
	PosFactor       PositionCompare
}

//ChronicFatigugeRankStatus is the info of the degree of fatigue
type ChronicFatigugeRankStatus struct {
	AnalysisPublic ChronicFatiguePublicInfo
	AnalysisData   []ChronicFatigueStatus `json:"analysis_data"`
}

//ChronicFatigueStatus is the member of ChronicFatigugeRankStatus
type ChronicFatigueStatus struct {
	Name       string  `json:"statue_name"`
	Percentage float64 `json:"percentage"`
}

//FatigueFactor is the member of ChronicFatigueData
type FatigueFactor struct {
	AnalysisPublic ChronicFatiguePublicInfo
	FatiFactorDtl  []FatigueFactorDetail
}

//FatigueFactorDetail is the info of the FatigueFactor
type FatigueFactorDetail struct {
	FactorName string `json:"factor_name"`
	FactorNum  int    `json:"factor_num"`
}

//SexCompare is 性别比较
type SexCompare struct {
	AnalysisPublic ChronicFatiguePublicInfo
	SexComInfo     []CompareInfo
}

//AgeCompare is 年龄比较
type AgeCompare struct {
	AnalysisPublic ChronicFatiguePublicInfo
	AgeFactorInfo  []AgeComInfo
}

//AgeComInfo is the member of AgeCompare
type AgeComInfo struct {
	FactorName   string `json:"factor_name"`
	AgeFactorNum []int  `json:"age_factor_num"`
	// Detail     AgeStaInfo `json:"detail"`
}

// //AgeStaInfo is the member of AgeComInfo
// type AgeStaInfo struct {
// 	AgeName      string `json:"age_name"`
// 	AgeFactorNum []int  `json:"age_factor_num"`
// }

//PositionCompare is 职位比较
type PositionCompare struct {
	AnalysisPublic ChronicFatiguePublicInfo
	PosComInfo     []CompareInfo
}

//CompareInfo is the member of SexCompare and
type CompareInfo struct {
	Name    string `json:"compare_name"`
	Num     int    `json:"blue_num"`
	NextNum int    `json:"red_num"`
}

//ChronicFatiguePublicInfo is
type ChronicFatiguePublicInfo struct {
	AnalysisName   string `json:"analysis_name"`
	AnalysisDesc   string `json:"analysis_desc"`
	AnalysisResult string `json:"analysis_result"`
}
