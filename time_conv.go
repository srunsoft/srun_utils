package utils

import (
	"time"
)

// TimeFromInt64 @title TimeFromInt64
// @description 时间戳 -> time.Time
// @auth 黄宇超 时间（2021/4/2）
func TimeFromInt64(n int64) time.Time {
	return time.Unix(n, 0)
}

// TimeToString @title TimeToString
// @description time.Time -> string， 格式为 YYYY-MM-DD hh:mm:ss
// @auth 黄宇超 时间（2021/4/2）
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// TimeFromString @title TimeFromString
// @description string -> time.Time
// @auth 黄宇超 时间（2021/4/2）
func TimeFromString(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
}

// TimeStampToTimeString @title TimeStampToTimeString
// @description 将时间戳转为可读的字符串
// @auth 黄宇超 时间（2021/4/2）
func TimeStampToTimeString(timestamp int64) string {
	return TimeToString(TimeFromInt64(timestamp))
}

// TimeStringToTimeStamp @title TimeStringToTimeStamp
// @description 将字符串转化为时间戳
// @auth 黄宇超 时间（2021/4/2）
func TimeStringToTimeStamp(timestr string) (timestamp int64, err error) {
	t, err := TimeFromString(timestr)
	if err != nil {
		return
	}
	timestamp = t.Unix()
	return
}

func NowString() string {
	return TimeToString(time.Now())
}

func Days(days int) int64 {
	return int64(days) * 24 * 3600
}

func Months(months int) int64 {
	return Days(30)
}

func Years(years int) int64 {
	return Days(365)
}

func TimeSubstract(t time.Time, dt int64) time.Time {
	return TimePlus(t, -dt)
}

func TimePlus(t time.Time, dt int64) time.Time {
	tmp := t.Unix() + dt
	return time.Unix(tmp, 0)
}