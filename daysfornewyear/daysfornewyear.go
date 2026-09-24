package daysfornewyear

import (
	"time"
)

func DaysUntilNewYear(t time.Time) int {
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	target := time.Date(today.Year() + 1, time.January, 1, 0, 0, 0, 0, today.Location())
	d := target.Sub(today)
	return int(d.Hours() / 24)
}