package cfas

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

type CountGroupOut struct {
	GroupID   string `json:"groupId"`
	GroupName string `json:"groupName"`
	GroupType int    `json:"groupType"`
	RegionID  string `json:"regionId"`
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
	GroupID string `json:"groupId"`

	FlowInNum  int     `json:"flowInNum"`
	FlowOutNum int     `json:"flowOutNum"`
	HoldValue  float64 `json:"holdValue"`

	NoRepeatInNum  int `json:"noRepeatInNum"`
	NoRepeatOutNum int `json:"noRepeatOutNum"`
	NetValue       int `json:"netValue"`
}
