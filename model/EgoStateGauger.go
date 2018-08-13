package model

//EgoState is the definition of The Ego-State Model test result struct
type EgoState struct {
	EgoStateChart []EgoStateDetail     `json:"ego_state_chart"`
	EgoStateInfo  []EgoStateInfoDetail `json:"ego_state_info"`
}

//EgoStateDetail is detials  of The EgoState  struct
type EgoStateDetail struct {
	Name          string `json:"name"`
	PositiveScore int    `json:"positive_score"`
	NegativeScore int    `json:"negative_score"`
}

//EgoStateInfoDetail is the struct of Ego-State Model classify
type EgoStateInfoDetail struct {
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	EgoState  []EgoStateDesc
}

//EgoStateDesc is the details of EgoStateInfoDetail struct
type EgoStateDesc struct {
	EgoStateName string `json:"egoState_name"`
	EgoDesc      string `json:"egoState_desc"`
	EgoDetail    string `json:"egoState_detail"`
}
