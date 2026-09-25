package daysfornewyear

import (
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	cases := []struct {
		name  string
		input time.Time
		want  int
	}{
		{"начало года", time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC), 365},
		{"конец года", time.Date(2026, time.December, 31, 0, 0, 0, 0, time.UTC), 1},
		{"високосный год", time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), 366},
		{"1 марта не високосного", time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC), 306},
		{"1 марта високосного", time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC), 306},
		{"28 февраля не високосного", time.Date(2026, time.February, 28, 0, 0, 0, 0, time.UTC), 307},
		{"28 февраля високосного", time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC), 308},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := DaysUntilNewYear(c.input)
			if got != c.want {
				t.Errorf("got %d; want %d", got, c.want)
			}
		})
	}
}
