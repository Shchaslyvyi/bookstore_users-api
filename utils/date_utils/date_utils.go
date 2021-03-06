package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

// GetNow func returns the current UTC time.
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString func returns the formatted current UTC time as string.
//func GetNowString() string {
//	return GetNow().Format(apiDateLayout)
//}

// GetNowDBFormat func returns the formatted current UTC time as DateTime string.
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
