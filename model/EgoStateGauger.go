package model

//EgoState is the definition of The Ego-State Model test result struct
type EgoState struct {
	EgoStateBrief EgoStateBriefInfo `json:"ego_state_brief"`
	EgoStateChart []EgoStateDetail  `json:"ego_state_chart"`
	EgoStateInfo  EgoStateClassfy   `json:"ego_state_info"`
}

//EgoStateBriefInfo is the memeber of EgoState
type EgoStateBriefInfo struct {
	BriefInfo     string         `json:"brief_info"`
	ClassifyBrief []EgoBriefInfo `json:"classify_brief"`
	ClassifyInfo  string         `json:"classify_info"`
}

//EgoBriefInfo is the member of EgoStateBriefInfo struct
type EgoBriefInfo struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

//EgoStateDetail is detials  of The EgoState  struct
type EgoStateDetail struct {
	Name          string `json:"name"`
	PositiveScore int    `json:"positive_score"`
	NegativeScore int    `json:"negative_score"`
}

//EgoStateClassfy is the member of EgoState struct
type EgoStateClassfy struct {
	ParentEgo EgoStateInfoDetail `json:"parent"`
	AdultEgo  EgoStateInfoDetail `json:"adult"`
	ChildEgo  EgoStateInfoDetail `json:"child"`
}

//EgoStateInfoDetail is the struct of Ego-State Model classify
type EgoStateInfoDetail struct {
	Name      string      `json:"name"`
	Introduce string      `json:"introduce"`
	EgoState  []EgoStates `json:"ego_state"`
}

//EgoStates is a member of EgoStateInfoDetail struct
type EgoStates struct {
	Name    string         `json:"ego_name"`
	Desc    string         `json:"ego_desc"`
	Details []EgoStateDesc `json:"details"`
}

//EgoStateDesc is the details of EgoStates struct
type EgoStateDesc struct {
	EgoStateName string `json:"egoState_name"`
	EgoDesc      string `json:"egoState_desc"`
	EgoDetail    string `json:"egoState_detail"`
}
