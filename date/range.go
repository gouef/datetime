package date

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
	RangeDateRegexp = `(` + datetime.YearRegexp + `-` + datetime.MonthRegexp + `-` + datetime.DayRegexp + `)`
	RangeRegexp     = `^([\[\(])` + RangeDateRegexp + `?\s*,\s*` + RangeDateRegexp + `?([\]\)])$`
)

type Range struct {
	from  Value
	to    Value
	start datetime.RangeStart
	end   datetime.RangeEnd
}

func NewRange(from, to string, start datetime.RangeStart, end datetime.RangeEnd) (*Range, error) {

	_, err := FromString(from)

	if from != "" && err != nil {
		return nil, err
	}

	_, err = FromString(to)

	if to != "" && err != nil {
		return nil, err
	}

	if to == "" && from == "" {
		return nil, errors.New("from and to can not be both empty")
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
	openBracket, date1, date2, closeBracket := match[1], match[2], match[6], match[10]

	return NewRange(date1, date2, datetime.RangeStart(openBracket), datetime.RangeEnd(closeBracket))
}

func (d *Range) Start() datetime.RangeStart {
	return d.start
}

func (d *Range) End() datetime.RangeEnd {
	return d.end
}

func (d *Range) From() Value {
	return d.from
}

func (d *Range) To() Value {
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

	from, _ := FromString(string(d.From()))
	to, _ := FromString(string(d.To()))
	return date.Between(from, to)
}

func (d *Range) format(date any) (datetime.Interface, error) {
	switch i := date.(type) {
	case time.Time:
		return New(i.Year(), int(i.Month()), i.Day())
	case *Date:
		return i, nil
	case Date:
		return &i, nil
	case string:
		return FromString(i)
	default:
		return nil, errors.New("unsupported format of date")
	}
}
