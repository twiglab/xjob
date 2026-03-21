package hdp

import "time"

func DayBeginEnd(t time.Time) (begin time.Time, end time.Time) {
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	next := today.Add(24 * time.Hour)
	return today, next
}

func Yestoday(now time.Time) time.Time {
	return now.Add(-1 * 24 * time.Hour)
}

func DT(t time.Time) string {
	return t.Format("20060102")
}
