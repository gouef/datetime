package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

type DateTimeValue string

func StringToDateTimeValue(value string) (DateTimeValue, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: dateTimeRegexp})

	if len(errs) != 0 {
		return "", errors.New(fmt.Sprintf("unsupported format of date time \"%s\"", value))
	}

	return DateTimeValue(value), nil
}

func (d DateTimeValue) Date() *DateTime {
	date, err := DateTimeFromString(string(d))

	if err != nil {
		return nil
	}

	return date
}
