package utils

import "time"

const (
	YYYY_MM_DD          = "2006-01-02"
	DATE_DIR_PATTERN    = "2006/01/02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
)

// FormatDateToString 将日期转换成指定格式的字符串
func FormatDateToString(date time.Time, format string) string {
	return date.Format(format)
}

// ParseStringToDate 将字符串日期转换成日期
func ParseStringToDate(date string, format string) (time.Time, error) {
	return time.Parse(format, date)
}

// TimeStrSub 日期字符串相减 datestr1-datestr2
func TimeStrSub(datestr1, datestr2, format string) (int64, error) {
	d1, err := ParseStringToDate(datestr1, format)
	if err != nil {
		return 0, err
	}

	d2, err := ParseStringToDate(datestr2, format)
	if err != nil {
		return 0, err
	}

	result := d1.Unix() - d2.Unix()

	return result, nil
}
