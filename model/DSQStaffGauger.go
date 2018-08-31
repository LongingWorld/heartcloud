package model

//DSQReportData 防御方式自评量表个人报告数据结构
type DSQReportData struct {
	DSQBrief   DSQBriefInfo  `json:"dsq_brief"`
	DSQSense   string        `json:"dsq_sense"`
	DSQResult  DSQResultData `json:"dsq_result"`
	DSQDeclare string        `json:"dsq_declare"`
}

//DSQBriefInfo 简介数据结构
type DSQBriefInfo struct {
	BriefInfo   DSQDetailInfo `json:"brief_info"`
	DSQClassify []DSQClassify `json:"dsq_classify"`
}

//DSQClassify 防御方式分类数据结构
type DSQClassify struct {
	ClassifyName   string          `json:"classify_name"`
	ClassifyInfo   string          `json:"classify_info"`
	ClassifyDetail []DSQDetailInfo `json:"classify_detail"`
}

//DSQResultData 测评结果数据结构
type DSQResultData struct {
	Concealment     DSQDetailInfo      `json:"concealment"`
	DSQTestInfo     DSQMechanismDetail `json:"result_info"`
	DSQMature       DSQMechanismDetail `json:"dsq_mature"`
	DSQIntermediate DSQMechanismDetail `json:"dsq_intermediate"`
	DSQNotMature    DSQMechanismDetail `json:"dsq_not_mature"`
}

//DSQMechanismDetail 防御因子得分说明
type DSQMechanismDetail struct {
	Explain         string           `json:"explain"`
	DSQFactorScores []DSQDetailScore `json:"dsq_scores"`
	DSQNote         string           `json:"dsq_suggest"`
}

//DSQDetailScore 防御因子得分数据结构
type DSQDetailScore struct {
	FactorName  string `json:"factor_name"`
	FactorScore int    `json:"factor_score"`
}

//DSQDetailInfo 因子说明及结果说明
type DSQDetailInfo struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
