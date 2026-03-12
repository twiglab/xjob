package cfas

// application/json
// Content-Type:text/plain;charset=UTF-8

type Rtn[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total int `json:"total"`
		List  []T `json:"list"`
	} `json:"data"`
}

// 1.15查询统计组数据
// ArtemisURL: /api/cfas/v2/countGroup/groups/page
type CountGroupIn struct {
	GroupType int    `json:"groupType"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
	Name      string `json:"name,omitempty"`
}

// 1.8查询时间范围内的多个统计组的客流统计数据
// ArtemisURL: /api/cfas/v2/passengerFlow/groups

type PassengerFlowIn struct {
	IDs         string `json:"ids"`
	Granularity string `json:"granularity"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type PassengerFlowOut struct {
	StartTime string `json:"startTime"`
	GroupID   string `json:"groupId"`
	GroupName string `json:"groupName"`

	FlowInNum  int `json:"flowInNum"`
	FlowOutNum int `json:"flowOutNum"`
	HoldValue  int `json:"holdValue"`

	NoRepeatInNum  int `json:"noRepeatInNum"`
	NoRepeatOutNum int `json:"noRepeatOutNum"`
	NetValue       int `json:"netValue"`
}
