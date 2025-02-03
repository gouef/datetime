package tests

import (
	"github.com/gouef/datetime"
	"github.com/gouef/datetime/date"
	"github.com/gouef/datetime/time"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValue(t *testing.T) {

	t.Run("DateValue", func(t *testing.T) {
		str := "2025-02-02"
		dateValue, err := date.StringToValue(str)

		assert.Nil(t, err)
		expected, _ := date.FromString(str)
		assert.Equal(t, expected, dateValue.Date())
	})

	t.Run("DateValue invalid", func(t *testing.T) {
		str := "2025-02-31"
		dateValue, err := date.StringToValue(str)
		timeValueDate := dateValue.Date()

		assert.Error(t, err)
		assert.Nil(t, timeValueDate)
	})

	t.Run("TimeValue", func(t *testing.T) {
		str := "18:30:05"
		timeValue, err := time.StringToValue(str)

		assert.Nil(t, err)
		expected, _ := time.FromString(str)
		assert.Equal(t, expected, timeValue.Date())
	})

	t.Run("TimeValue invalid", func(t *testing.T) {
		str := "18:60:05"
		timeValue, err := time.StringToValue(str)
		timeValueDate := timeValue.Date()

		assert.Error(t, err)
		assert.Nil(t, timeValueDate)
	})

	t.Run("Value", func(t *testing.T) {
		str := "2025-02-02 18:30:05"
		dateTimeValue, err := datetime.StringToValue(str)

		assert.Nil(t, err)
		expected, _ := datetime.FromString(str)
		assert.Equal(t, expected, dateTimeValue.Date())
	})

	t.Run("Value invalid", func(t *testing.T) {
		str := "2025-02-31 18:30:05"
		dateTimeValue, err := datetime.StringToValue(str)
		timeValueDate := dateTimeValue.Date()

		assert.Error(t, err)
		assert.Nil(t, timeValueDate)
	})

}
