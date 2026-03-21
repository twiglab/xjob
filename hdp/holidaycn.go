package hdp

import (
	"encoding/json/v2"
	"os"
	"time"
)

type Holiday struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	IsOffDay bool   `json:"isOffDay"`
	NotFound bool   `json:"-"`
}

type Holidays struct {
	Year   int       `json:"year"`
	Papers []string  `json:"papers"`
	Days   []Holiday `json:"days"`
}

func (h Holidays) Find(t time.Time) Holiday {
	s := t.Format(time.DateOnly)

	for _, d := range h.Days {
		if d.Date == s {
			return d
		}
	}

	return Holiday{NotFound: true}
}

func LoadFile(jsfile string) (Holidays, error) {
	var hs Holidays
	f, err := os.Open(jsfile)
	if err != nil {
		return hs, err
	}
	defer f.Close()

	err = json.UnmarshalRead(f, &hs)
	return hs, err
}
