package wanf

const (
	SizeHour  = "hour"
	SizeDay   = "day"
	SizeWeek  = "week"
	SizeMonth = "month"
	SizeYear  = "year"
)

const (
	IndexPassBy          = "passBy"
	IndexEnter           = "enter"
	IndexExits           = "exits"
	IndexTurnBack        = "trunBack"
	IndexDuplicatePeople = "duplicatePeople"
)

func Index(s ...string) []string {
	return s
}

type PeopleCountingIn struct {
	Types int      `json:"types"`
	Site  []int    `json:"site"`
	Size  string   `json:"size"`
	Index []string `json:"index"`
	Date  string   `json:"data"`
}

type PeopleCountingOut struct {
	Code int `json:"code"`
}

type PeopleCoutingItem struct {
	Enter int `json:"enter__sum"`
	Exits int `json:"exits__sum"`
}
