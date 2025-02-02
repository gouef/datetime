package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

type TimeValue string

func StringToTimeValue(value string) (TimeValue, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: timeRegexp})

	if len(errs) != 0 {
		return "", errors.New(fmt.Sprintf("unsupported format of time \"%s\"", value))
	}

	return TimeValue(value), nil
}

func (d TimeValue) Date() *Time {
	date, err := TimeFromString(string(d))

	if err != nil {
		return nil
	}

	return date
}
