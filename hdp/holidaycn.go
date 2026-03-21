package hdp

type Holiday struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	IsOffDay bool   `json:"isOffDay"`
}

type Holidays struct {
	Year   int      `json:"year"`
	Papers []string `json:"papers"`
	Days   Holiday  `json:"days"`
}

func Load(jsfile ...string) []Holidays {
	return nil
}
