package model

//DSQCompanyReportData 防御方式企业报告数据结构
type DSQCompanyReportData struct {
	DSQComBrief           string             `json:"dsq_com_brief"`
	DSQComDataAnalysis    DSQComDataAnalysis `json:"dsq_com_data_analysis"`
	DSQComResultSummarize string             `json:"dsq_com_result_summarize"`
	DSQComReportSuggest   string             `json:"dsq_com_suggest"`
}

//DSQComDataAnalysis 防御方式企业报告数据分析数据结构
type DSQComDataAnalysis struct {
	DSQConcealment DSQComDetailInfo      `json:"dsq_concealment"`
	DSQHealthState DSQComHealthStateData `json:"dsq_health_data"`
	DSQKindsDetail DSQComKindsDetail     `json:"dsq_kinds_detail"`
}

//DSQComHealthStateData 防御机制健康状况数据结构
type DSQComHealthStateData struct {
	Name         string          `json:"dsq_health_name"`
	Info         string          `json:"dsq_health_info"`
	Data         []DSQComPersent `json:"dsq_health_details"`
	DataAnalysis string          `json:"dsq_health_annlysis"`
}

//DSQComKindsDetail 各类防御机制具体情况数据结构
type DSQComKindsDetail struct {
	DSQName             string            `json:"dsq_kinds_name"`
	DSQMatureInfo       DSQComDataDetails `json:"dsq_mature_name"`
	DSQIntermediateInfo DSQComDataDetails `json:"dsq_intermediate_name"`
	DSQNotMatureInfo    DSQComDataDetails `json:"dsq_not_mature_name"`
}

//DSQComDataDetails 各类防御机制具体情况-详细信息
type DSQComDataDetails struct {
	KindsName         string             `json:"dsq_classify_name"`
	KindsInfo         string             `json:"dsq_class_info"`
	MaturePersent     DSQComPersentDTL   `json:"dsq_mature_persent"`
	TotalDataAanlysis DSQComDetailInfo   `json:"dsq_total_analysis"`
	FactorInfo        string             `json:"dsq_factor_info"`
	FactorDetails     []DSQComDetailInfo `json:"dsq_factor_details"`
	FactorPerName     string             `json:"dsq_staff_use"`
	FactorPerDetail   []DSQComPersent    `json:"dsq_persent_factor"`
	FactorAnalysis    string             `json:"dsq_factor_analysis"`
}

//DSQComPersent  防御因子百分比数据结构
type DSQComPersent struct {
	FactorName    string `json:"factor_name"`
	FactorPersent DSQComPersentDTL
}

//DSQComPersentDTL 百分比数据
type DSQComPersentDTL struct {
	CommonPersent float64 `json:"common_persent"`
	OftenPersent  float64 `json:"often_persent"`
	OncePresent   float64 `json:"once_persent"`
}

//DSQComDetailInfo 因子说明及结果说明
type DSQComDetailInfo struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
