package datetime

import (
	"time"
)

const DateYYYYMMDDhhmmssLayout = "2006-01-02 15:04:05"
const DateYYYYMMDDLayout = "2006-01-02"

// NowDateTime 获取现在的时间
func NowDateTime() string {
	return time.Now().Format(DateYYYYMMDDhhmmssLayout)
}

// NowDateYYYYMMDD 获取现在的日期
func NowDateYYYYMMDD() string {
	return time.Now().Format(DateYYYYMMDDLayout)
}

// GetUnixTimeByYYYYMMDDhhmmss 指定时间转时间戳
func GetUnixTimeByYYYYMMDDhhmmss(timeStr string) (time.Time, error) {
	return time.ParseInLocation(DateYYYYMMDDhhmmssLayout, timeStr, time.Local)
}
func GetUnixTimeByYYYYMMDD(timeStr string) (time.Time, error) {
	return time.ParseInLocation(DateYYYYMMDDLayout, timeStr, time.Local)
}

func GetDateTimeByUnix(unix int64) string {
	if unix == 0 {
		return time.Unix(time.Now().Unix(), 0).Format(DateYYYYMMDDhhmmssLayout)
	}
	return time.Unix(unix, 0).Format(DateYYYYMMDDhhmmssLayout)
}
func GetDateByUnix(unix int64) string {
	if unix == 0 {
		return time.Unix(time.Now().Unix(), 0).Format(DateYYYYMMDDLayout)
	}
	return time.Unix(unix, 0).Format(DateYYYYMMDDLayout)
}
