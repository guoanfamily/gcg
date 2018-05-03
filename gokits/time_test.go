package utils

import "testing"

func TestToTime(t *testing.T) {
	t.Logf("%s to time: %v", "2017-09-24", ToTime("yyyy-MM-dd", "2017-09-14"))
}
