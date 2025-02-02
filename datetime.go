package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/utils"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"strconv"
	"time"
)

const (
	Regexp = `^(\d{4})-(\d{2})-(\d{2})( (\d{2}):(\d{2}):(\d{2}))?$`
)

type DateTime struct {
	Year     int
	Month    int `validate:"min=1,max=12"`
	Day      int `validate:"min=1,max=31"`
	Hour     int `validate:"min=0,max=23"`
	Minute   int `validate:"min=0,max=59"`
	Second   int `validate:"min=0,max=59"`
	DateTime time.Time
}

func Now() *DateTime {
	now := time.Now()

	return &DateTime{
		Year:     now.Year(),
		Month:    int(now.Month()),
		Day:      now.Day(),
		Hour:     now.Hour(),
		Minute:   now.Minute(),
		Second:   now.Second(),
		DateTime: now,
	}
}

func New(year, month, day, hour, minute, second int) (*DateTime, error) {
	errs := validator.Validate(year, constraints.GreaterOrEqual{Value: 0})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("year must be 0 or greater get \"%d\"", year))
	}

	errs = validator.Validate(month, constraints.Range{Min: 1, Max: 12})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("month must be between 1-12 get \"%d\"", month))
	}
	daysInMonth := DaysInMonth(year, month)
	errs = validator.Validate(day, constraints.Range{Min: 1, Max: float64(daysInMonth)})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("day must be between 1-%d for month %d of year %d get \"%d\"", daysInMonth, month, year, day))
	}

	errs = validator.Validate(hour, constraints.Range{Min: 0, Max: 23})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("hour must be between 0-23 get \"%d\"", hour))
	}

	errs = validator.Validate(minute, constraints.Range{Min: 0, Max: 59})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("minute must be between 0-59 get \"%d\"", minute))
	}

	errs = validator.Validate(second, constraints.Range{Min: 0, Max: 59})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("second must be between 0-59 get \"%d\"", second))
	}

	return &DateTime{
		Year:     year,
		Month:    month,
		Day:      day,
		Hour:     hour,
		Minute:   minute,
		Second:   second,
		DateTime: time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC),
	}, nil
}

func FromString(value string) (Interface, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: Regexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date \"%s\"", value))
	}

	re := regexp.MustCompile(Regexp)
	match := re.FindStringSubmatch(value)
	year, _ := strconv.Atoi(match[1])
	month, _ := strconv.Atoi(match[2])
	day, _ := strconv.Atoi(match[3])

	if match[4] != "" {
		hour, _ := strconv.Atoi(match[5])
		minute, _ := strconv.Atoi(match[6])
		second, _ := strconv.Atoi(match[7])

		return New(year, month, day, hour, minute, second)
	}

	return nil, errors.New(fmt.Sprintf("unsupported format of datetime \"%s\". time is missing.", value))
}

func (d *DateTime) FromString(value string) (Interface, error) {
	return FromString(value)
}

func (d *DateTime) ToString() string {
	return d.Time().Format(time.DateTime)
}

func GetDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (d *DateTime) IsWeekend() bool {
	weekendDays := []time.Weekday{time.Sunday, time.Saturday}
	return utils.InArray(d.DateTime.Weekday(), weekendDays)
}

func (d *DateTime) Time() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

// Compare compares the date instant d with u. If d is before u, it returns -1;
// if d is after u, it returns +1; if they're the same, it returns 0.
func (d *DateTime) Compare(u Interface) int {
	return d.Time().Compare(u.Time())
}

func (d *DateTime) Equal(u Interface) bool {
	return d.Time().Equal(u.Time())
}

func (d *DateTime) Between(start, end Interface) bool {
	return d.Time().Before(end.Time()) && d.Time().After(start.Time())
}

func DaysInMonthList(year int, month int) []int {
	days := make([]int, DaysInMonth(year, month))

	for i := range days {
		days[i] = i + 1
	}
	return days
}

func DaysInMonth(year int, month int) int {
	return DaysInMonthByDate(GetDate(year, month, 1))
}

func DaysInMonthByDate(t time.Time) int {
	y, m, _ := t.Date()
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
