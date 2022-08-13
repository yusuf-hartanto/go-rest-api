package utils

import (
	"time"
)

func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func GetPeriode(month string) (periodes []string) {
	start, _ := time.Parse("2006-01-02", month+"-01")
	end, _ := time.Parse("2006-01-02", month+"-30")
	for rd := RangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		periodes = append(periodes, date.String()[0:10])
	}
	return periodes
}

func GetFirstAndLastDayofMonth(month string) (time.Time, time.Time) {
	date, _ := time.Parse("2006-01", month)
	y, m, _ := date.Date()
	loc := date.Location()

	firstDay := time.Date(y, m, 1, 0, 0, 0, 0, loc)
	lastDay := time.Date(y, m+1, 1, 0, 0, 0, -1, loc)
	return firstDay, lastDay
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
