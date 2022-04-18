package utils

import "time"

// 定义常用的默认格式
const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
)

// 声明时区
var (
	location    = time.UTC                      // 国际标准时
	locationCST = time.FixedZone("GMT", 8*3600) // 东八(防止win无时区文件)
)

var (
	monthLayout  = TIME_LAYOUT[:7] // 年月
	dayLayout    = TIME_LAYOUT[:10]
	hourLayout   = TIME_LAYOUT[:13]
	minuteLayout = TIME_LAYOUT[:16] // 年月日时分
	fullLayout   = TIME_LAYOUT
)

// 指定时间变换时区
func LocChange(t time.Time, loc *time.Location) time.Time {
	if loc == nil { // 不设定,则使用标准时,对标Parse将时间解释为UTC时间
		loc = location
	}

	return t.In(loc)
}

// 指定时间的当天开始时间
func DayStart(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 指定时间的当天结束时间
func DayEnd(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 1e9-1, t.Location())
}

// 指定时间的当天开始/结束时间
func DayRange(t time.Time) (time.Time, time.Time) {
	return DayStart(t), DayEnd(t)
}

// 指定时间的当周开始时间, 即周一00:00:00
func WeekStart(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}

	weekDay := t.Weekday()
	day := t.Day()

	switch weekDay {
	case time.Sunday: // Sunday 值为0,
		day -= 6
	default:
		day -= int(weekDay) - 1
	}

	return time.Date(t.Year(), t.Month(), day, 0, 0, 0, 0, t.Location())
}

// 指定时间的当周结束时间, 即周日23:59:59
func WeekEnd(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}

	weekDay := t.Weekday()
	day := t.Day()
	switch weekDay {
	case time.Sunday: // 周日不用处理

	default: // 其天时间，比如周1 需要加 7-1==6 天
		day += 7 - int(weekDay)
	}

	return time.Date(t.Year(), t.Month(), day, 23, 59, 59, 1e9-1, t.Location())
}

// 指定时间的当周开始/结束时间 即:周一00:00和周日23:59:59
func WeekRange(t time.Time) (time.Time, time.Time) {
	return WeekStart(t), WeekEnd(t)
}

// 指定时间的当月开始时间, 即:当月第一天00:00:00
func MonthStart(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}

	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// 指定时间的当月结束时间, 即:当月最后一天23:59:59
func MonthEnd(t time.Time) time.Time {
	if t.IsZero() {
		return time.Now()
	}

	month := t.Month()
	day := 31

	switch month {
	case time.February: // 2 月
		if year := t.Year(); year%400 == 0 || (year%100 != 0 && year%4 == 0) { // 闰年29天
			day = 29
		} else { // 非闰年28天
			day = 28
		}

	case 4, 6, 9, 11: // 小月 30天
		day = 30
	}

	return time.Date(t.Year(), t.Month(), day, 23, 59, 59, 1e9-1, t.Location())
}

// 指定时间的当月开始/结束时间
func MonthRange(t time.Time) (time.Time, time.Time) {
	return MonthStart(t), MonthEnd(t)
}

// 指定二个时间是否是同一天
func SameDay(first, last time.Time) bool {
	return time.Time(first).Year() == time.Time(last).Year() && time.Time(first).YearDay() == time.Time(last).YearDay()
}

// DiffNano 时间差，纳秒
func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}
