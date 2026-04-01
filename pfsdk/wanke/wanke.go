package wanke

type PeopleCountingIn struct {
	Types int   `json:"types"`
	Site  []int `json:"site"`
	Size  string
	Index []string
	Date  string
}

type PeopleCountingOut struct {
}
