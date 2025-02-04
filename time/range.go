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
	RangeRegexp = `^([\[\(])(\d{4}-\d{2}-\d{2})?,(\d{4}-\d{2}-\d{2})?([\]\)])$`
)

type Range struct {
	from  Value
	to    Value
	start datetime.RangeStart
	end   datetime.RangeEnd
}

func NewRange(from, to string, start datetime.RangeStart, end datetime.RangeEnd) *Range {
	return &Range{
		from:  Value(from),
		to:    Value(to),
		start: start,
		end:   end,
	}
}

func NewRangeOptional(from, to string) *Range {
	return NewRange(from, to, datetime.RangeStartOptional, datetime.RangeEndOptional)
}

func NewRangeStrict(from, to string) *Range {
	return NewRange(from, to, datetime.RangeStartStrict, datetime.RangeEndStrict)
}

func NewRangeStartStrict(from, to string) *Range {
	return NewRange(from, to, datetime.RangeStartStrict, datetime.RangeEndOptional)
}

func NewRangeStartOptional(from, to string) *Range {
	return NewRange(from, to, datetime.RangeStartOptional, datetime.RangeEndStrict)
}

func RangeFromString(dateRange string) (*Range, error) {
	errs := validator.Validate(dateRange, constraints.RegularExpression{Regexp: RangeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date range \"%s\"", dateRange))
	}

	re := regexp.MustCompile(RangeRegexp)
	match := re.FindStringSubmatch(dateRange)
	openBracket, date1, date2, closeBracket := match[1], match[2], match[3], match[4]

	return NewRange(date1, date2, datetime.RangeStart(openBracket), datetime.RangeEnd(closeBracket)), nil
}

func (d *Range) Start() datetime.RangeStart {
	return d.start
}

func (d *Range) End() datetime.RangeEnd {
	return d.end
}

func (d *Range) String() string {
	return fmt.Sprintf("%s%s,%s%s", d.start, d.from, d.to, d.end)
}

func (d *Range) Is(value any) bool {
	date, err := d.format(value)

	if err != nil {
		return false
	}

	start, _ := FromString(string(d.start))
	end, _ := FromString(string(d.end))
	return date.Between(start, end)
}

func (d *Range) format(date any) (datetime.Interface, error) {
	switch i := date.(type) {
	case time.Time:
		return New(i.Year(), int(i.Month()), i.Day())
	case *Time:
		return i, nil
	case Time:
		return &i, nil
	case string:
		return FromString(i)
	default:
		return nil, errors.New("unsupported format of date")
	}
}
