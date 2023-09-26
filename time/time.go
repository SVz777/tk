package time

import (
	"errors"
	"time"

	"github.com/SVz777/tk/convert"
)

const (
	DefaultTime     = "1971-01-01 00:00:00"
	LayoutTime      = "2006-01-02 15:04:05"
	LayoutTimeYMDHM = "2006-01-02 15:04"
	LayoutTimeYMDH  = "2006-01-02 15"
	LayoutTimeYMD   = "2006-01-02"
	LayoutTimeYM    = "2006-01"
	LayoutTimeYmd   = "20060102"
	LayoutTimeYm    = "200601"
	LayoutTimeY     = "2006"
	LayoutTimeHI    = "15:04"

	ChineseLayoutTime = "2006年1月2日 15:04:05"
	ChineseLayoutMD   = "01月02日"
)

// GetCurrentDateTime 获取当前时间 YYYY-MM-DD H:i:s
func GetCurrentDateTime() (currentTime string) {
	return time.Now().Format(LayoutTime)
}

// GetCurrentChineseDateTime 获取当前中文时间 YYYY年MM月DD日 H:i:s
func GetCurrentChineseDateTime() (currentTime string) {
	return time.Now().Format(ChineseLayoutTime)
}

// TimestampToDate 时间戳转日期
func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(LayoutTime)
}

// TimestampToDateWithLayout 时间戳转换为日期，输出指定格式
func TimestampToDateWithLayout(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// GetDayDateTimeWithExtra 获取指定的年月日 YYYY年MM月DD日 ，+/- time
func GetDayDateTimeWithExtra(extra string) string {
	curTime := time.Now() // 获取系统当前时间
	dh, _ := time.ParseDuration(extra)
	return curTime.Add(dh).Format(LayoutTimeYMD)
}

// GetDayDateTimeWithExtraLayout 获取指定的年月日 YYYY年MM月DD日 ，+/- time
func GetDayDateTimeWithExtraLayout(extra string, layout string, date string) (string, error) {
	afterDate, err := time.Parse(LayoutTime, date)
	if err != nil {
		return "", err
	}
	dh, _ := time.ParseDuration(extra)
	return afterDate.Add(dh).Format(layout), nil
}

// GetDayStart 获取开始日期
func GetDayStart(date string, extra string) (result string, err error) {
	datetime, err := time.Parse(LayoutTime, date)
	if nil != err {
		return "", err
	}
	targetDatetime := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())
	dh, _ := time.ParseDuration(extra)
	return targetDatetime.Add(dh).Format(LayoutTime), nil
}

// GetDayEnd 获取结束日期
func GetDayEnd(date string, extra string) (result string, err error) {
	datetime, err := time.Parse(LayoutTime, date)
	if nil != err {
		return "", err
	}
	targetDatetime := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 23, 59, 59, 0, datetime.Location())
	dh, _ := time.ParseDuration(extra)
	return targetDatetime.Add(dh).Format(LayoutTime), nil
}

// DatetimeToTimestamp 日期 转 时间戳
func DatetimeToTimestamp(date string, layout string) (timestamp int64, err error) {
	if layout == "" {
		layout = LayoutTime
	}
	datetime, err := time.Parse(layout, date)
	if nil != err {
		return 0, err
	}
	return datetime.Unix(), nil
}

// TimestampToDatetime 时间戳 转 日期
func TimestampToDatetime(timestamp int64, layout string) (datetime string, err error) {
	if layout == "" {
		layout = LayoutTime
	}
	t := time.Unix(timestamp, 0)
	return t.Format(layout), nil
}

var defaultParseTimeLayouts = []string{
	LayoutTime,
	LayoutTimeYMDH,
	LayoutTimeYMDHM,
	LayoutTimeYMD,
	LayoutTimeYM,
	LayoutTimeY,
}

// ParseTime 日期解析
func ParseTime(dateTime string, layouts ...string) (time.Time, error) {
	if len(layouts) == 0 {
		layouts = defaultParseTimeLayouts
	}

	for idx := range layouts {
		if t, err := time.ParseInLocation(layouts[idx], dateTime, time.Local); err == nil {
			return t, nil
		}
	}

	return time.Now(), errors.New("parsing time error:" + dateTime)
}

// GetMonthDays 获取月份天数
func GetMonthDays(year, month int) int {
	d1 := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	d2 := d1.AddDate(0, 1, -1)
	return d2.Day() - d1.Day() + 1
}

// GetYearMonth 获取年月 eg.202002
func GetYearMonth(yearMonthS string) int32 {
	d, _ := ParseTime(yearMonthS)
	ymS := d.Format(LayoutTimeYm)
	ym, _ := convert.Int32(ymS)
	return ym
}

// GetYearMonthDay 获取年月日 eg.20200202
func GetYearMonthDay(yearMonthdayS string) int32 {
	d, _ := ParseTime(yearMonthdayS)
	ymdS := d.Format(LayoutTimeYmd)
	ymd, _ := convert.Int32(ymdS)
	return ymd
}

// GetYearMonthString 获取年月字符串
func GetYearMonthString(yearMonth int32) string {
	ym, _ := convert.String(yearMonth)
	d, _ := time.ParseInLocation(LayoutTimeYm, ym, time.Local)
	return d.Format(LayoutTimeYM)
}
