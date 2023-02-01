package utils

import (
	"strconv"
	"time"
)

const DateLayout string = "2006-01-02"
const DateTimeLayout string = "2006-01-02 15:04:05"
const DateTimeWithMicroLayout string = "2006-01-02 15:04:05.000000"
const DateTimeWithoutDashLayout = "20060102150405"

var (
	China, _          = time.LoadLocation("Asia/Chongqing")
	FirstDayOf2020, _ = time.ParseInLocation(DateTimeLayout, "2020-01-01 00:00:00", China)
)

func NormalizeDate(t time.Time) time.Time {
	s := t.Format(DateLayout)
	t, _ = time.Parse(DateLayout, s)
	return t
}

func FirstDayOfMonth(t time.Time) time.Time {
	return NormalizeDate(t.AddDate(0, 0, -t.Day()+1))
}

func ParseDate(dateString string, defaultDate time.Time) time.Time {
	if date, err := Date(dateString); nil != err {
		return defaultDate
	} else {
		return date
	}
}

func DatetimeWithMicroLayout(t string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation(DateTimeWithMicroLayout, t, location)
}

func Date(date string) (time.Time, error) {
	return time.Parse(DateLayout, date)
}

func Datetime(t string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation(DateTimeLayout, t, location)
}

func ToDatetimeString(t time.Time) string {
	return t.Format(DateTimeLayout)
}

func ToDatetimeStringWithoutDash(t time.Time) string {
	return t.Format(DateTimeWithoutDashLayout)
}

func NowInChinaToStringWithMicroLayout() string {
	return time.Now().In(China).Format(DateTimeWithMicroLayout)
}

func ToDateString(t time.Time) string {
	return t.Format(DateLayout)
}

const maxTimestamp = uint64(99999999999999)

func ReverseTimeStamp(timestamp string) string {
	value, _ := strconv.ParseUint(timestamp, 10, 64)
	return strconv.FormatUint(maxTimestamp-value, 10)
}
