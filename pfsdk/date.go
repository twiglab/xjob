package pfsdk

import "time"

func OpenTime(t time.Time) (start, end, eight time.Time) {
	start = time.Date(t.Year(), t.Month(), t.Day(), 10, 0, 0, 0, t.Location())
	end = time.Date(t.Year(), t.Month(), t.Day(), 22, 0, 0, 0, t.Location())
	eight = time.Date(t.Year(), t.Month(), t.Day(), 20, 0, 0, 0, t.Location())
	return
}

func Yestoday(now time.Time) time.Time {
	return now.Add(-1 * 24 * time.Hour)
}

func DateOnly(t time.Time) string {
	return t.Format(time.DateOnly)
}

func DT(t time.Time) string {
	return t.Format("20060102")
}
