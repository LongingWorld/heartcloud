package model

//EgoCompanyData 自我状态企业量表报告数据
type EgoCompanyData struct {
	BriefInfo           EgoCompanyBriefInfo       `json:"brief_info"`
	ResultAnalysis      EgoCompanyResultAnalysis  `json:"result_analysis"`
	QualityRiskAnalysis EgoComQualityRiskAnalysis `json:"quality_risk_analysis"`
}

//EgoCompanyBriefInfo 自我状态企业报告简介数据
type EgoCompanyBriefInfo struct {
	Info     string               `json:"info"`
	Classify []EgoCompanyClassify `json:"classify"`
	Desc     string               `json:"desc"`
}

//EgoCompanyClassify 自我状态分类信息数据
type EgoCompanyClassify struct {
	ClassifyName string `json:"classify_name"`
	ClassifyDesc string `json:"classify_desc"`
}

//EgoCompanyResultAnalysis 自我状态企业报告结果分析数据
type EgoCompanyResultAnalysis struct {
	ResultInfo string                 `json:"result_info"`
	ResultData []EgoCompanyResultData `json:"result_data"`
}

//EgoCompanyResultData 自我状态企业报告结果分析详细信息
type EgoCompanyResultData struct {
	EgoDimName     string             `json:"ego_dim_name"`
	EgoDimData     EgoComResultDetail `json:"ego_dim_data"`
	EgoDimAnalysis string             `json:"ego_dim_analysis"`
}

//EgoComResultDetail 自我状态企业报告结果分析详细信息
type EgoComResultDetail struct {
	SuperExcellent float64 `json:"super_excellent"`
	Excellent      float64 `json:"excellent "`
	Medium         float64 `json:"medium"`
	Danger         float64 `json:"danger"`
	HighDanger     float64 `json:"high_danger"`
}

//EgoComQualityRiskAnalysis 自我状态企业报告-管理团队管理品质与风险分析
type EgoComQualityRiskAnalysis struct {
	Info         string                  `json:"risk_ana_info"`
	AnalysisData []EgoComQualityRiskInfo `json:"analysis_data"`
	AnalysisInfo string                  `json:"analysis_info"`
	Suggestion   string                  `json:"suggestion"`
}

//EgoComQualityRiskInfo 自我状态企业报告-管理团队管理品质与风险详细数据
type EgoComQualityRiskInfo struct {
	QualityRiskName string `json:"quality_risk_name"`
	QualityRiskNum  int    `json:"quality_risk_num"`
}
