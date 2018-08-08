package model

/*ChronicFatigueStaffReport is the response of report_company template_id=4*/
type ChronicFatigueStaffReport struct {
	Section1 ChronicFatigueSection1 `json:"Section1"`
	Section2 ChronicFatigueSection2 `json:"Section2"`
	//Section3 ChronicFatigueSection3 `json:"Section3"`
	//TemplateID int                    `json:"template_id"`
}

/*ChronicFatigueSection1 is*/
type ChronicFatigueSection1 struct {
	Introduction string                       `json:"introduction"`
	Classify     []ChronicFatigueNormDescribe `json:"classify"`
}

/*ChronicFatigueSection2 is*/
type ChronicFatigueSection2 struct {
	AccordItemNum int                           `json:"accord_item_num"`
	AccordExplain string                        `json:"accord_explain"`
	DimensionInfo []ChronicFatigueDimensionDesc `json:"dimension_info"`
	//ChronicFatigueEndanger []ChronicFatigueNormDescribe  `json:"chronic_fatigue_endanger"`
}

// /*ChronicFatigueSection3 is*/
// type ChronicFatigueSection3 struct {
// 	SuggestDesc string                       `json:"suggest_desc"`
// 	Suggests    []ChronicFatigueNormDescribe `json:"suggests"`
// }

/*ChronicFatigueNormDescribe is */
type ChronicFatigueNormDescribe struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

/*ChronicFatigueDimensionDesc is*/
type ChronicFatigueDimensionDesc struct {
	Name         string   `json:"name"`
	IsInclude    int      `json:"is_include"`
	DimDesc      string   `json:"dim_desc"`
	SuggestDesc  string   `json:"suggest_desc"`
	SubjectNames []string `json:"field"`
}
