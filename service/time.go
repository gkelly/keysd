package service

import (
	"time"
)

// DateDisplay returns date string representation of a time value
func dateDisplay(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	today := time.Now()
	if today.Format("01/02/2006") == t.Format("01/02/2006") {
		return "today"
	}
	yesterday := today.AddDate(0, 0, -1)
	if yesterday.Format("01/02/2006") == t.Format("01/02/2006") {
		return "yesterday"
	}
	return t.Format("01/02")
}

// TimeDisplay returns time string representation of a time value
func timeDisplay(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("3:04pm")
}
