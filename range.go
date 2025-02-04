package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"time"
)

type Range struct {
	from  Value
	to    Value
	start RangeStart
	end   RangeEnd
}

func NewRange(from, to string, start RangeStart, end RangeEnd) (*Range, error) {

	_, err := FromString(from)

	if from != "" && err != nil {
		return nil, err
	}

	_, err = FromString(to)

	if to != "" && err != nil {
		return nil, err
	}

	return &Range{
		from:  Value(from),
		to:    Value(to),
		start: start,
		end:   end,
	}, nil
}

func NewRangeOptional(from, to string) (*Range, error) {
	return NewRange(from, to, RangeStartOptional, RangeEndOptional)
}

func NewRangeStrict(from, to string) (*Range, error) {
	return NewRange(from, to, RangeStartStrict, RangeEndStrict)
}

func NewRangeStartStrict(from, to string) (*Range, error) {
	return NewRange(from, to, RangeStartStrict, RangeEndOptional)
}

func NewRangeStartOptional(from, to string) (*Range, error) {
	return NewRange(from, to, RangeStartOptional, RangeEndStrict)
}

func RangeFromString(value string) (*Range, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: RangeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of datetime range \"%s\"", value))
	}

	re := regexp.MustCompile(RangeRegexp)
	match := re.FindStringSubmatch(value)
	start := match[1]
	from := match[2]
	to := match[9]
	end := match[16]

	return NewRange(from, to, RangeStart(start), RangeEnd(end))
}

func (r *Range) Start() RangeStart {
	return r.start
}

func (r *Range) End() RangeEnd {
	return r.end
}

func (r *Range) From() ValueInterface {
	return r.from
}

func (r *Range) To() ValueInterface {
	return r.to
}

func (r *Range) String() string {
	return fmt.Sprintf("%s%s, %s%s", r.Start(), r.From(), r.To(), r.End())
}

func (r *Range) Is(value any) bool {
	date, err := r.format(value)

	if err != nil {
		return false
	}

	from, _ := FromString(string(r.from))
	to, _ := FromString(string(r.to))

	if from == nil && to == nil {
		return false
	}

	if from == nil {
		return date.Before(to)
	}

	if to == nil {
		return date.After(from)
	}

	return date.Between(from, to)
}

func (r *Range) format(date any) (Interface, error) {
	switch i := date.(type) {
	case time.Time:
		return New(i.Year(), int(i.Month()), i.Day(), i.Hour(), i.Minute(), i.Second())
	case *DateTime:
		return i, nil
	case string:
		return FromString(i)
	default:
		return nil, errors.New("unsupported format of datetime")
	}
}
