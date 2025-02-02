package tests

import (
	"github.com/gouef/datetime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValue(t *testing.T) {

	t.Run("DateValue", func(t *testing.T) {
		str := "2025-02-02"
		dateValue, err := datetime.StringToDateValue(str)

		assert.Nil(t, err)
		expected, _ := datetime.DateFromString(str)
		assert.Equal(t, expected, dateValue.Date())
	})

	t.Run("TimeValue", func(t *testing.T) {
		str := "18:30:05"
		timeValue, err := datetime.StringToTimeValue(str)

		assert.Nil(t, err)
		expected, _ := datetime.TimeFromString(str)
		assert.Equal(t, expected, timeValue.Date())
	})

	t.Run("TimeValue invalid", func(t *testing.T) {
		str := "18:60:05"
		timeValue, err := datetime.StringToTimeValue(str)
		timeValueDate := timeValue.Date()

		assert.Nil(t, err)
		assert.Nil(t, timeValueDate)
	})

	t.Run("DateTimeValue", func(t *testing.T) {
		str := "2025-02-02 18:30:05"
		dateTimeValue, err := datetime.StringToDateTimeValue(str)

		assert.Nil(t, err)
		expected, _ := datetime.DateTimeFromString(str)
		assert.Equal(t, expected, dateTimeValue.Date())
	})

}
