package datetime

import (
	"time"
)

var TIME_LAYOUT string = "2006-01-02 15:04:05"

func TimeToString(t time.Time) string {
	return t.Format(TIME_LAYOUT)
}

func StringToTime(s string) (time.Time, error) {
	t, err := time.Parse(TIME_LAYOUT, s)
	return t, err
}
