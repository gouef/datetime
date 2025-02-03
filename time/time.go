package time

import (
	"errors"
	"fmt"
	"github.com/gouef/datetime"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"strconv"
	goTime "time"
)

const (
	Regexp         = `^(\d{2}):(\d{2}):(\d{2})?$`
	DateTimeRegexp = `^(((\d{4})-(\d{2})-(\d{2}))( ))?((\d{2}):(\d{2}):(\d{2}))$`
)

type Time struct {
	Hour     int `validate:"min=0,max=23"`
	Minute   int `validate:"min=0,max=59"`
	Second   int `validate:"min=0,max=59"`
	DateTime goTime.Time
}

func New(hour, minute, second int) (datetime.Interface, error) {
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
		DateTime: goTime.Date(0, goTime.Month(1), 1, hour, minute, second, 0, goTime.UTC),
	}, nil
}

func FromString(value string) (datetime.Interface, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: DateTimeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date \"%s\"", value))
	}

	re := regexp.MustCompile(DateTimeRegexp)
	match := re.FindStringSubmatch(value)
	hour, _ := strconv.Atoi(match[8])
	minute, _ := strconv.Atoi(match[9])
	second, _ := strconv.Atoi(match[10])

	return New(hour, minute, second)
}

func (t *Time) FromString(value string) (datetime.Interface, error) {
	return FromString(value)
}

func (t *Time) ToString() string {
	return t.Time().Format(goTime.TimeOnly)
}

func (t *Time) Time() goTime.Time {
	return goTime.Date(0, goTime.Month(1), 1, t.Hour, t.Minute, t.Second, 0, goTime.UTC)
}

// Compare compares the date instant d with u. If d is before u, it returns -1;
// if d is after u, it returns +1; if they're the same, it returns 0.
func (t *Time) Compare(u datetime.Interface) int {
	return t.Time().Compare(u.Time())
}

func (t *Time) Equal(u datetime.Interface) bool {
	return t.Time().Equal(u.Time())
}

func (t *Time) Between(start, end datetime.Interface) bool {
	return t.Time().Before(end.Time()) && t.Time().After(start.Time())
}
