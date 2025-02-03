package datetime

import (
	"errors"
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

type Value string

func StringToValue(value string) (Value, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: Regexp})

	d, err := FromString(value)

	if len(errs) != 0 || err != nil {
		return "", errors.New(fmt.Sprintf("unsupported format of date time \"%s\"", value))
	}

	str := d.ToString()

	return Value(str), nil
}

func (d Value) Date() Interface {
	date, err := FromString(string(d))

	if err != nil {
		return nil
	}

	return date
}
