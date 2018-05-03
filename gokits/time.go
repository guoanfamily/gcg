package utils

import (
	"fmt"
	"strings"
	"time"
)

// GetLocalTime get local time
func GetLocalTime() *time.Time {
	now := time.Now()
	return &now
}

// GetUTCTime get utc time
func GetUTCTime() *time.Time {
	now := time.Now().UTC()
	return &now
}

// GetCurrMilliTs 获取当前时间戳
func GetCurrMilliTs() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetZeroTime 获取当天0点时间
func GetZeroTime() time.Time {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	zeroTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return zeroTime
}

// GetZeroTimeOf 获取指定日期的0点时间
func GetZeroTimeOf(t time.Time) time.Time {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	zeroTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return zeroTime
}

// GetNextDayZeroTs 获取下一天的0点时间戳
func GetNextDayZeroTs() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day() + 1
	zeroTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return zeroTime.Unix()
}

// 获取指定时间戳的时间
func GetTimeByTs(ts int64) time.Time {
	time := time.Unix(ts, 0)
	return time
}

// TsToCn 获取时间(60s->1分钟，60分钟->1小时)
func TsToCn(ts int32) string {
	minutes := ts / 60
	var hour int32 = 0
	var minute int32 = 0
	str := ""
	if minutes >= 60 {
		hour = minutes / 60
		minute = minutes % 60
		if minute == 0 {
			str = fmt.Sprintf("%d小时", hour)
		} else {
			str = fmt.Sprintf("%d小时%d分钟", hour, minute)
		}
	} else {
		str = fmt.Sprintf("%d分钟", minutes)
	}

	return str
}

// TimeAgoInWords 多久之前
func TimeAgoInWords(ts int64) string {
	now := time.Now().Unix()
	minutes := int32(0)
	timeAgoInWords := ""
	if ts > 0 && now > ts {
		minutes = int32((now - ts) / 60)

		if minutes >= 0 && minutes <= 59 {
			if minutes == 0 {
				timeAgoInWords = "刚刚"
			} else {
				timeAgoInWords = fmt.Sprintf("%d分钟前", minutes)
			}
		} else if minutes >= 60 && minutes <= 1439 {
			timeAgoInWords = fmt.Sprintf("%d小时前", minutes/60)
		} else if minutes >= 1440 && minutes <= 10079 {
			timeAgoInWords = fmt.Sprintf("%d天前", minutes/1440)
		} else if minutes >= 10080 && minutes <= 43199 {
			timeAgoInWords = fmt.Sprintf("%d星期前", minutes/11520)
		} else if minutes >= 43200 && minutes <= 525599 {
			timeAgoInWords = fmt.Sprintf("%d个月前", minutes/43200)
		} else {
			timeAgoInWords = fmt.Sprintf("%d年前", minutes/525600)
		}
	}

	return timeAgoInWords
}

// TimeFormat 日期格式化: yyyy-MM-dd HH:mm:ss
func TimeFormat(layout string, dateTime time.Time) string {
	converts := map[string]string{
		"yyyy": "2006",
		"MM":   "01",
		"dd":   "02",
		"HH":   "15",
		"mm":   "04",
		"ss":   "05",
	}
	for k, v := range converts {
		layout = strings.Replace(layout, k, v, 1)
	}
	return dateTime.Format(layout)
}

// ToTime 时间字符串转 Time
func ToTime(layout, value string) *time.Time {
	converts := map[string]string{
		"yyyy": "2006",
		"MM":   "01",
		"dd":   "02",
		"HH":   "15",
		"mm":   "04",
		"ss":   "05",
	}
	for k, v := range converts {
		layout = strings.Replace(layout, k, v, 1)
	}

	dateTime, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		panic(err)
	}
	return &dateTime
}
