package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

type DateValue string

func StringToDateValue(value string) (DateValue, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: dateRegexp})

	if len(errs) != 0 {
		return "", errors.New(fmt.Sprintf("unsupported format of date \"%s\"", value))
	}

	return DateValue(value), nil
}

func (d DateValue) Date() *Date {
	date, err := DateFromString(string(d))

	if err != nil {
		return nil
	}

	return date
}
