package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"regexp"
	"time"
)

const (
	dateRangeRegexp = `^([\[\(])(\d{4}-\d{2}-\d{2})?,(\d{4}-\d{2}-\d{2})?([\]\)])$`
)

type DateRange struct {
	from  string
	to    string
	start RangeStart
	end   RangeEnd
}

func NewDateRange(from, to string, start RangeStart, end RangeEnd) *DateRange {
	return &DateRange{
		from:  from,
		to:    to,
		start: start,
		end:   end,
	}
}

func (d *DateRange) String() string {
	return fmt.Sprintf("%s%s,%s%s", d.start, d.from, d.to, d.end)
}

func (d *DateRange) Is(value any) bool {
	date, err := d.format(value)

	if err != nil {
		return false
	}

	start, _ := DateFromString(string(d.start))
	end, _ := DateFromString(string(d.end))
	return date.Between(start, end)
}

func DateRangeFromString(dateRange string) (*DateRange, error) {
	errs := validator.Validate(dateRange, constraints.RegularExpression{Regexp: dateRangeRegexp})

	if len(errs) != 0 {
		return nil, errors.New(fmt.Sprintf("unsupported format of date range \"%s\"", dateRange))
	}

	re := regexp.MustCompile(dateRangeRegexp)
	match := re.FindStringSubmatch(dateRange)
	openBracket, date1, date2, closeBracket := match[1], match[2], match[3], match[4]

	return NewDateRange(date1, date2, RangeStart(openBracket), RangeEnd(closeBracket)), nil
}

func (d *DateRange) format(date any) (*Date, error) {
	switch i := date.(type) {
	case time.Time:
		return NewDate(i.Year(), int(i.Month()), i.Day())
	case *Date:
		return i, nil
	case Date:
		return &i, nil
	case string:
		return DateFromString(i)
	default:
		return nil, errors.New("unsupported format of date")
	}
}
