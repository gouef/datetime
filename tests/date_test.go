package tests

import (
	"fmt"
	"github.com/gouef/datetime"
	"github.com/gouef/datetime/date"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		day      int
		expected *date.Date
		err      bool
	}{
		{2024, 12, 25, &date.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, false},
		{2024, 2, 30, nil, true},   // Invalid day for February
		{2024, 13, 10, nil, true},  // Invalid month (13)
		{2024, 4, 31, nil, true},   // April has only 30 days
		{-2024, 12, 25, nil, true}, // April has only 30 days
	}

	for _, tt := range tests {
		t.Run("TestNewDate", func(t *testing.T) {
			d, err := date.New(tt.year, tt.month, tt.day)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, d)
			}
		})
	}
}

func TestDateIsWeekend(t *testing.T) {
	tests := []struct {
		date     *date.Date
		expected bool
	}{
		{&date.Date{Year: 2024, Month: 12, Day: 21, DateTime: time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC)}, true},  // Saturday
		{&date.Date{Year: 2024, Month: 12, Day: 22, DateTime: time.Date(2024, 12, 22, 0, 0, 0, 0, time.UTC)}, true},  // Sunday
		{&date.Date{Year: 2024, Month: 12, Day: 23, DateTime: time.Date(2024, 12, 23, 0, 0, 0, 0, time.UTC)}, false}, // Monday
	}

	for _, tt := range tests {
		t.Run("TestIsWeekend", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date.IsWeekend())
		})
	}
}

func TestDateCompare(t *testing.T) {
	tests := []struct {
		date1    *date.Date
		date2    *date.Date
		expected int
	}{
		{&date.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)},
			&date.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, 0},
		{&date.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)},
			&date.Date{Year: 2024, Month: 12, Day: 26, DateTime: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC)}, -1}, // 25th < 26th
		{&date.Date{Year: 2024, Month: 12, Day: 26, DateTime: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC)},
			&date.Date{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, 1}, // 26th > 25th
	}

	for _, tt := range tests {
		t.Run("TestCompare", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date1.Compare(tt.date2))
		})
	}
}

func TestDateEqual(t *testing.T) {
	tests := []struct {
		date1    *date.Date
		date2    *date.Date
		expected bool
	}{
		// Test 1: Equal DateTime instances
		{
			date1:    &date.Date{Year: 2024, Month: 3, Day: 31},
			date2:    &date.Date{Year: 2024, Month: 3, Day: 31},
			expected: true,
		},
		// Test 2: Different DateTime instances (different day)
		{
			date1:    &date.Date{Year: 2024, Month: 3, Day: 31},
			date2:    &date.Date{Year: 2024, Month: 3, Day: 30},
			expected: false,
		},
		// Test 3: Different DateTime instances (different month)
		{
			date1:    &date.Date{Year: 2024, Month: 3, Day: 31},
			date2:    &date.Date{Year: 2024, Month: 4, Day: 1},
			expected: false,
		},
		// Test 4: Different DateTime instances (different year)
		{
			date1:    &date.Date{Year: 2024, Month: 3, Day: 31},
			date2:    &date.Date{Year: 2025, Month: 3, Day: 31},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run("TestEqual", func(t *testing.T) {
			result := tt.date1.Equal(tt.date2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDateBetween(t *testing.T) {
	date1, _ := date.New(2025, 2, 1)
	date2, _ := date.New(2024, 2, 1)
	date3, _ := date.New(2026, 2, 1)

	tests := []struct {
		date     datetime.Interface
		start    datetime.Interface
		end      datetime.Interface
		expected bool
	}{
		{date1, date2, date3, true},
		{date2, date1, date3, false},
		{date3, date2, date3, false},
	}

	for _, tt := range tests {
		t.Run("TestDaysInMonth", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date.Between(tt.start, tt.end))
		})
	}
}

func TestDateFromString(t *testing.T) {
	validDate, _ := date.New(2025, 1, 31)
	tests := []struct {
		date         string
		expectedErr  bool
		expectedDate datetime.Interface
	}{
		{"2025-01-31", false, validDate},
		{"2025-01-31 23:27:15", false, validDate},
		{"2025-02-31", true, nil},
		{"2025-13-32", true, nil},
		{"-2025-06-31", true, nil},
		{"invalid", true, nil},
	}

	for _, tt := range tests {
		t.Run("TestDateFromString: "+tt.date, func(t *testing.T) {
			d, err := date.FromString(tt.date)
			if tt.expectedErr {
				assert.Nil(t, d)
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedDate, d)

				d2, err := tt.expectedDate.FromString(tt.date)
				assert.Nil(t, err)
				assert.Equal(t, d2, d)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	validDate, _ := date.New(2025, 1, 31)
	tests := []struct {
		expectedString string
		date           string
		expectedErr    bool
		expectedDate   datetime.Interface
	}{
		{"2025-01-31", "2025-01-31", false, validDate},
		{"2025-01-31", "2025-01-31 23:27:15", false, validDate},
		{"2025-02-31", "2025-02-31", true, nil},
		{"2025-13-32", "2025-13-32", true, nil},
		{"-2025-06-31", "-2025-06-31", true, nil},
		{"invalid", "invalid", true, nil},
	}

	for _, tt := range tests {
		t.Run("TestDateToString: "+tt.date, func(t *testing.T) {
			d, err := date.FromString(tt.date)
			if tt.expectedErr {
				assert.Nil(t, d)
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedDate, d)
				assert.Equal(t, tt.expectedString, d.ToString())
			}
		})
	}
}

func TestDateRange(t *testing.T) {
	t.Run("New Date Range", func(t *testing.T) {
		tests := []struct {
			rangeStr         string
			testDate         any
			expectedErr      bool
			expectedValidErr bool
			expected         func() any
		}{
			{"[2025-01-31, 2026-01-31]", "2025-05-06", false, false, func() any {
				val, err := date.NewRangeStrict(
					"2025-01-31", "2026-01-31")
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", true, false, func() any {
				val, err := date.NewRange(
					"2025-01-31 14:15:16", "2026-01-31 17:18:19", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", true, false, func() any {
				val, err := date.NewRange(
					"", "2026-01-31 17:18:19", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, ]", "2025-05-06 20:21:22", true, false, func() any {
				val, err := date.NewRange(
					"2025-01-31 14:15:16", "", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, ]", time.Date(2025, 5, 6, 20, 21, 22, 0, time.UTC), true, false, func() any {
				val, err := date.NewRange(
					"2025-01-31 14:15:16", "", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, 2026-01-31]", "2025-05-06", true, false, nil},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", 2025, true, true, func() any {
				val, err := date.NewRange(
					"2025-01-31 14:15:16", "2026-01-31 17:18:19", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[, ]", date.Now(), true, true, func() any {
				val, err := date.NewRange(
					"", "", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.Error(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, 2026-02-31 14:15:16]", date.Now(), true, true, func() any {
				val, err := date.NewRange(
					"2025-01-31 14:15:16", "2026-02-31 14:15:16", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.Error(t, err)
				return val
			}},
		}

		for _, tt := range tests {
			t.Run("New Range: "+tt.rangeStr+fmt.Sprintf(", valid date: %v", tt.testDate), func(t *testing.T) {
				if tt.expectedErr {
					r, err := date.RangeFromString(tt.rangeStr)
					assert.Error(t, err)
					assert.Nil(t, r)

				} else {
					r, err := date.RangeFromString(tt.rangeStr)
					assert.NoError(t, err)
					assert.Equal(t, tt.expected(), r)
					if tt.expectedValidErr {
						assert.False(t, r.Is(tt.testDate))
					} else {
						assert.True(t, r.Is(tt.testDate))
						assert.Equal(t, tt.rangeStr, r.String())
					}
				}
			})
		}

	})

	t.Run("Range Is", func(t *testing.T) {
		tests := []struct {
			rangeStr    string
			testDate    any
			expectedErr bool
			expected    bool
		}{
			{"[2025-01-31, 2026-01-31]", "2025-05-06", false, true},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", true, true},
			{"[, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", true, true},
			{"[2025-01-31 14:15:16, ]", "2025-05-06 20:21:22", true, true},
			{"[2025-01-31 14:15:16, ]", time.Date(2025, 5, 6, 20, 21, 22, 0, time.UTC), true, true},
			{"[2025-01-31 14:15:16, 2026-01-31]", "2025-05-06", true, false},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", 2025, true, false},
			{"[, ]", date.Now(), true, false},
			{"[2, ]", date.Now(), true, false},
		}

		for _, tt := range tests {
			t.Run("New Range: "+tt.rangeStr+fmt.Sprintf(", valid date: %v", tt.testDate), func(t *testing.T) {
				if tt.expectedErr {
					r, err := date.RangeFromString(tt.rangeStr)
					assert.Error(t, err)
					assert.Nil(t, r)

				} else {
					r, err := date.RangeFromString(tt.rangeStr)
					assert.NoError(t, err)
					assert.Equal(t, tt.expected, r.Is(tt.testDate))
				}
			})
		}
	})

	t.Run("Range brackets", func(t *testing.T) {
		tests := []struct {
			expectedErr bool
			start       string
			from        string
			to          string
			end         string
			expected    func() any
		}{
			{false, "[", "2025-01-31", "2026-01-31", "]", func() any {
				val, err := date.NewRangeStrict("2025-01-31", "2026-01-31")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "2026-01-31 17:18:19", "]", func() any {
				val, err := date.NewRangeStrict("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "", "2026-01-31 17:18:19", "]", func() any {
				val, err := date.NewRangeStrict("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "", "]", func() any {
				val, err := date.NewRangeStrict("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "2026-01-31 17:18:19", "]", func() any {
				val, err := date.NewRangeStartOptional("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "", "2026-01-31 17:18:19", "]", func() any {
				val, err := date.NewRangeStartOptional("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "", "]", func() any {
				val, err := date.NewRangeStartOptional("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "2026-01-31 17:18:19", ")", func() any {
				val, err := date.NewRangeOptional("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "", "2026-01-31 17:18:19", ")", func() any {
				val, err := date.NewRangeOptional("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "", ")", func() any {
				val, err := date.NewRangeOptional("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "2026-01-31 17:18:19", ")", func() any {
				val, err := date.NewRangeStartStrict("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "", "2026-01-31 17:18:19", ")", func() any {
				val, err := date.NewRangeStartStrict("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "", ")", func() any {
				val, err := date.NewRangeStartStrict("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
		}

		for _, tt := range tests {
			t.Run("New Range brackets: "+fmt.Sprintf("From: %s%s, To: %s%s", tt.start, tt.from, tt.to, tt.end), func(t *testing.T) {
				r, err := date.NewRange(tt.from, tt.to, datetime.RangeStart(tt.start), datetime.RangeEnd(tt.end))

				if tt.expectedErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.expected(), r)
				}
			})
		}
	})
}
