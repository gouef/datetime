package time

import (
	"errors"
	"fmt"
	"github.com/gouef/datetime"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

type Value string

func StringToValue(value string) (Value, error) {
	errs := validator.Validate(value, constraints.RegularExpression{Regexp: DateTimeRegexp})

	time, err := FromString(value)

	if len(errs) != 0 || err != nil {
		return "", errors.New(fmt.Sprintf("unsupported format of time \"%s\"", value))
	}

	str := time.ToString()

	return Value(str), nil
}

func (d Value) String() string {
	return string(d)
}

func (d Value) Date() datetime.Interface {
	time, err := FromString(d.String())

	if err != nil {
		return nil
	}

	return time
}
