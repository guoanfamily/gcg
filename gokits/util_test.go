package utils

import (
	"testing"
	"time"
)

func TestIsEmpty(t *testing.T) {
	t.Logf("string: a, %v", IsEmpty("a"))
	t.Logf("string: , %v", IsEmpty(""))
	t.Logf("int: 0, %v", IsEmpty(int(0)))
	t.Logf("int: 1, %v", IsEmpty(int(1)))
	t.Logf("int32: 0, %v", IsEmpty(int32(0)))
	t.Logf("int32: 1, %v", IsEmpty(int32(1)))
	t.Logf("int64: 0, %v", IsEmpty(int64(0)))
	t.Logf("int64: 1, %v", IsEmpty(int64(1)))
	t.Logf("byte: 0, %v", IsEmpty(byte(0)))
	t.Logf("byte: 1, %v", IsEmpty(byte(1)))

}

func TestParseQueries(t *testing.T) {
	t.Logf("%v", ParseQueries("https://uland.taobao.com/coupon/edetail?e=DHRskzIKen%2B7QfSPPGeGabpmFv4aEcr40UDGfotqzCKa5rBMT5rEqug5OxjP%2FnOttqrNb7%2Fe7nk9txhkdMJomB0HgBdG%2FDDL%2F1M%2FBw7Sf%2FfPd%2BzgVG4oJ%2FmjsDkW9dQWGzeLItWH3j9isbZR9N21PnbuxLy55xxw&pid=mm_96847743_12778058_116792905&af=1"))
}

func TestFormatNumberToShortString(t *testing.T) {
	testArgs := []int64{100, 10000, 10101, 10000000, 1000000001, 10000000000000}
	for _, arg := range testArgs {
		t.Logf("format %d to short string: %s", arg, FormatNumberToShortString(arg))
	}
}

func TestSleepRandomTime(t *testing.T) {
	maxSleepTime := time.Duration(10) * time.Second
	t.Logf("sleep random time, max sleep time: %fs", maxSleepTime.Seconds())
	beforeSleep := time.Now()
	SleepRandomTime(maxSleepTime)
	t.Logf("sleep time: %fs", time.Now().Sub(beforeSleep).Seconds())
}

func TestGenRandomString(t *testing.T) {
	randomRange := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	n := 4
	t.Logf("gen random string: %s", GenRandomString(n, randomRange))
}
