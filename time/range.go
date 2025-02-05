package time

import (
	"errors"
	"fmt"
	"github.com/gouef/datetime"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"time"
)

const (
	RangeRegexp = `^([\[\(])` + DateTimeRegexp + `?\s*,\s*` + DateTimeRegexp + `?([\]\)])$`
)

type Range struct {
	from  Value
	to    Value
	start datetime.RangeStart
	end   datetime.RangeEnd
}

func NewRange(from, to string, start datetime.RangeStart, end datetime.RangeEnd) (*Range, error) {

	if from == "" && to == "" {
		return nil, errors.New("from and to (both) can not be empty")
	}

	validFrom, err := getTimeFromDateTime(from)

	if err != nil && from != "" {
		return nil, errors.New(fmt.Sprintf("from (%s) is not valid", from))
	}

	validTo, err := getTimeFromDateTime(to)

	if err != nil && to != "" {
		return nil, errors.New(fmt.Sprintf("to (%s) is not valid", to))
	}

	if validFrom == nil {
		from = ""
	} else {
		from = validFrom.ToString()
	}

	if validTo == nil {
		to = ""
	} else {
		to = validTo.ToString()
	}

	return &Range{
		from:  Value(from),
		to:    Value(to),
		start: start,
		end:   end,
	}, nil
}

func NewRangeOptional(from, to string) (*Range, error) {
	return NewRange(from, to, datetime.RangeStartOptional, datetime.RangeEndOptional)
}

func NewRangeStrict(from, to string) (*Range, error) {
	return NewRange(from, to, datetime.RangeStartStrict, datetime.RangeEndStrict)
}

func NewRangeStartStrict(from, to string) (*Range, error) {
	return NewRange(from, to, datetime.RangeStartStrict, datetime.RangeEndOptional)
}

func NewRangeStartOptional(from, to string) (*Range, error) {
	return NewRange(from, to, datetime.RangeStartOptional, datetime.RangeEndStrict)
}

func RangeFromString(dateRange string) (*Range, error) {
	errs := validator.Validate(dateRange, constraints.RegularExpression{Regexp: RangeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date range \"%s\"", dateRange))
	}

	re := regexp.MustCompile(RangeRegexp)
	match := re.FindStringSubmatch(dateRange)
	openBracket, date1, date2, closeBracket := match[1], match[7], match[17], match[22]

	if date1 == "" && date2 == "" {
		return nil, errors.New(fmt.Sprintf("time range from and to can not be both empty \"%s\"", dateRange))
	}

	return NewRange(date1, date2, datetime.RangeStart(openBracket), datetime.RangeEnd(closeBracket))
}

func (d *Range) Start() datetime.RangeStart {
	return d.start
}

func (d *Range) End() datetime.RangeEnd {
	return d.end
}

func (d *Range) From() datetime.ValueInterface {
	return d.from
}

func (d *Range) To() datetime.ValueInterface {
	return d.to
}

func (d *Range) String() string {
	return fmt.Sprintf("%s%s, %s%s", d.Start(), d.From(), d.To(), d.End())
}

func (d *Range) Is(value any) bool {
	date, err := d.format(value)

	if err != nil {
		return false
	}

	from, _ := FromString(string(d.from))
	to, _ := FromString(string(d.to))

	if from == nil {
		return date.Before(to)
	}

	if to == nil {
		return date.After(from)
	}

	return date.Between(from, to)
}

func getTimeFromDateTime(date string) (datetime.Interface, error) {
	return FromString(date)
}

func (d *Range) format(date any) (datetime.Interface, error) {
	switch i := date.(type) {
	case time.Time:
		return New(i.Hour(), i.Minute(), i.Second())
	case *Time:
		return i, nil
	case string:
		return FromString(i)
	default:
		return nil, errors.New("unsupported format of date")
	}
}
