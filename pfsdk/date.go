package pfsdk

import "time"

func Yestoday(now time.Time) time.Time {
	return now.Add(-1 * 24 * time.Hour)
}

func DateOnly(t time.Time) string {
	return t.Format(time.DateOnly)
}

func DT(t time.Time) string {
	return t.Format("20060102")
}
