package tests

import (
	"github.com/gouef/datetime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		day      int
		expected *datetime.Date
		err      bool
	}{
		{2024, 12, 25, &datetime.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, false},
		{2024, 2, 30, nil, true},  // Invalid day for February
		{2024, 13, 10, nil, true}, // Invalid month (13)
		{2024, 4, 31, nil, true},  // April has only 30 days
	}

	for _, tt := range tests {
		t.Run("TestNewDate", func(t *testing.T) {
			date, err := datetime.NewDate(tt.year, tt.month, tt.day)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, date)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		date     *datetime.Date
		expected bool
	}{
		{&datetime.Date{Year: 2024, Month: 12, Day: 21, DateTime: time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC)}, true},  // Saturday
		{&datetime.Date{Year: 2024, Month: 12, Day: 22, DateTime: time.Date(2024, 12, 22, 0, 0, 0, 0, time.UTC)}, true},  // Sunday
		{&datetime.Date{Year: 2024, Month: 12, Day: 23, DateTime: time.Date(2024, 12, 23, 0, 0, 0, 0, time.UTC)}, false}, // Monday
	}

	for _, tt := range tests {
		t.Run("TestIsWeekend", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date.IsWeekend())
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		date1    datetime.Date
		date2    datetime.Date
		expected int
	}{
		{datetime.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)},
			datetime.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, 0},
		{datetime.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)},
			datetime.Date{Year: 2024, Month: 12, Day: 26, DateTime: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC)}, -1}, // 25th < 26th
		{datetime.Date{Year: 2024, Month: 12, Day: 26, DateTime: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC)},
			datetime.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, 1}, // 26th > 25th
	}

	for _, tt := range tests {
		t.Run("TestCompare", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date1.Compare(tt.date2))
		})
	}
}

func TestDaysInMonthList(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		expected []int
	}{
		{2024, 2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},          // Leap year February
		{2024, 4, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},      // April
		{2024, 12, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}}, // December
	}

	for _, tt := range tests {
		t.Run("TestDaysInMonthList", func(t *testing.T) {
			assert.Equal(t, tt.expected, datetime.DaysInMonthList(tt.year, tt.month))
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		expected int
	}{
		{2024, 2, 29},  // Leap year
		{2024, 4, 30},  // April
		{2024, 12, 31}, // December
	}

	for _, tt := range tests {
		t.Run("TestDaysInMonth", func(t *testing.T) {
			assert.Equal(t, tt.expected, datetime.DaysInMonth(tt.year, tt.month))
		})
	}
}

func TestDaysInMonthByDate(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected int
	}{
		{time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), 29},  // February in a leap year
		{time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), 30},  // April
		{time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC), 31}, // December
	}

	for _, tt := range tests {
		t.Run("TestDaysInMonthByDate", func(t *testing.T) {
			assert.Equal(t, tt.expected, datetime.DaysInMonthByDate(tt.date))
		})
	}
}
