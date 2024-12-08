package helpers

import "time"

func ParseTime(s string) time.Time {
	parsed, _ := time.Parse("15:04:05", s)
	return parsed
}