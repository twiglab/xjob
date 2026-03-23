package pfsdk

import "time"

type KeyTime struct {
	Zero      time.Time
	OpenStart time.Time
	OpenEnd   time.Time
	Eight     time.Time
}

func MakeKeyTime(t time.Time) (kt KeyTime) {
	kt.Zero = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	kt.OpenStart = time.Date(t.Year(), t.Month(), t.Day(), 10, 0, 0, 0, t.Location())
	kt.OpenEnd = time.Date(t.Year(), t.Month(), t.Day(), 22, 0, 0, 0, t.Location())
	kt.Eight = time.Date(t.Year(), t.Month(), t.Day(), 20, 0, 0, 0, t.Location())
	return
}

func Yestoday(now time.Time) time.Time {
	return now.Add(-1 * 24 * time.Hour)
}

func DateOnly(t time.Time) string {
	return t.Format(time.DateOnly)
}

func DateTime(t time.Time) string {
	return t.Format(time.DateTime)
}

func DT(t time.Time) string {
	return t.Format("20060102")
}

func MinPerDay(t time.Time) int {
	h, m, _ := t.Clock()
	return h*60 + m
}
