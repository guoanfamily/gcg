package utils

import "time"

func String(s string) *string {
	return &s
}

func Bool(i bool) *bool {
	return &i
}

func Byte(i byte) *byte {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Float64(i float64) *float64 {
	return &i
}

func Time(t time.Time) *time.Time {
	return &t
}
