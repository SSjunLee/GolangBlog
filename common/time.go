package common

import (
	"time"
)

func IsInTimeRange(t time.Time, timeRange time.Duration) bool {
	now := time.Now()
	if t.After(now) && time.Duration(t.Sub(now).Minutes()) < timeRange {
		return true
	}
	return false
}
