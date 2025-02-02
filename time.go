package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"strconv"
	"time"
)

const (
	timeRegexp = `^(\d{2}):(\d{2}):(\d{2})?$`
)

type Time struct {
	Hour     int `validate:"min=0,max=23"`
	Minute   int `validate:"min=0,max=59"`
	Second   int `validate:"min=0,max=59"`
	DateTime time.Time
}

func NewTime(hour, minute, second int) (*Time, error) {
	errs := validator.Validate(hour, constraints.Range{Min: 0, Max: 23})

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

	return &Time{
		Hour:     hour,
		Minute:   minute,
		Second:   second,
		DateTime: time.Date(0, time.Month(0), 0, hour, minute, second, 0, time.UTC),
	}, nil
}

func TimeFromString(value string) (*Time, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: dateTimeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date \"%s\"", value))
	}

	re := regexp.MustCompile(dateTimeRegexp)
	match := re.FindStringSubmatch(value)
	hour, _ := strconv.Atoi(match[1])
	minute, _ := strconv.Atoi(match[2])
	second, _ := strconv.Atoi(match[3])

	return NewTime(hour, minute, second)
}

func (d *Time) Time() time.Time {
	return time.Date(0, time.Month(0), 0, d.Hour, d.Minute, d.Second, 0, time.UTC)
}

// Compare compares the date instant d with u. If d is before u, it returns -1;
// if d is after u, it returns +1; if they're the same, it returns 0.
func (d *Time) Compare(u *Time) int {
	return d.Time().Compare(u.Time())
}

func (d *Time) Equal(u *Time) bool {
	return d.Time().Equal(u.Time())
}

func (d *Time) Between(start, end *Time) bool {
	return d.Time().Before(end.Time()) && d.Time().After(start.Time())
}
