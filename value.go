package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

type Value string

func StringToDateTimeValue(value string) (Value, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: Regexp})

	if len(errs) != 0 {
		return "", errors.New(fmt.Sprintf("unsupported format of date time \"%s\"", value))
	}

	return Value(value), nil
}

func (d Value) Date() Interface {
	date, err := FromString(string(d))

	if err != nil {
		return nil
	}

	return date
}
