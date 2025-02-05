package date

import (
	"errors"
	"fmt"
	"github.com/gouef/datetime"
	"github.com/gouef/utils"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"strconv"
	"time"
)

const (
	Regexp         = `^(\d{4})-(\d{2})-(\d{2})?$`
	DateTimeRegexp = `^(((\d{4})-(\d{2})-(\d{2}))( )?)((\d{2}):(\d{2}):(\d{2}))?$`
)

type Date struct {
	Year     int
	Month    int `validate:"min=1,max=12"`
	Day      int `validate:"min=1,max=31"`
	DateTime time.Time
}

func Now() *Date {
	now := time.Now()

	return &Date{
		Year:     now.Year(),
		Month:    int(now.Month()),
		Day:      now.Day(),
		DateTime: now,
	}
}

func New(year, month, day int) (datetime.Interface, error) {
	errs := validator.Validate(year, constraints.GreaterOrEqual{Value: 0})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("year must be 0 or greater get \"%d\"", year))
	}

	errs = validator.Validate(month, constraints.Range{Min: 1, Max: 12})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("month must be between 1-12 get \"%d\"", month))
	}
	daysInMonth := datetime.DaysInMonth(year, month)
	errs = validator.Validate(day, constraints.Range{Min: 1, Max: float64(daysInMonth)})

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("day must be between 1-%d for month %d of year %d get \"%d\"", daysInMonth, month, year, day))
	}

	return &Date{
		Year:     year,
		Month:    month,
		Day:      day,
		DateTime: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}, nil
}

func FromString(value string) (datetime.Interface, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: DateTimeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date \"%s\"", value))
	}

	re := regexp.MustCompile(DateTimeRegexp)
	match := re.FindStringSubmatch(value)
	year, _ := strconv.Atoi(match[3])
	month, _ := strconv.Atoi(match[4])
	day, _ := strconv.Atoi(match[5])

	return New(year, month, day)
}

func (d *Date) FromString(value string) (datetime.Interface, error) {
	return FromString(value)
}

func (d *Date) ToString() string {
	return d.Time().Format(time.DateOnly)
}

func (d *Date) IsWeekend() bool {
	weekendDays := []time.Weekday{time.Sunday, time.Saturday}
	return utils.InArray(d.DateTime.Weekday(), weekendDays)
}

func (d *Date) Time() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

// Compare compares the date instant d with u. If d is before u, it returns -1;
// if d is after u, it returns +1; if they're the same, it returns 0.
func (d *Date) Compare(u datetime.Interface) int {
	return d.Time().Compare(u.Time())
}

func (d *Date) Equal(u datetime.Interface) bool {
	return d.Time().Equal(u.Time())
}

func (d *Date) Between(start, end datetime.Interface) bool {
	return d.Before(end) && d.After(start)
}

func (d *Date) Before(u datetime.Interface) bool {
	return d.Time().Before(u.Time())
}

func (d *Date) After(u datetime.Interface) bool {
	return d.Time().After(u.Time())
}
