package tests

import (
	"fmt"
	"github.com/gouef/datetime"
	"github.com/gouef/datetime/date"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewDateTime(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		day      int
		hour     int
		minute   int
		second   int
		expected *datetime.DateTime
		err      bool
	}{
		{2024, 12, 25, 0, 0, 0, &datetime.DateTime{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, false},
		{2024, 2, 30, 0, 0, 0, nil, true},
		{2024, 13, 10, 0, 0, 0, nil, true},
		{2024, 4, 31, 0, 0, 0, nil, true},
		{-2024, 12, 25, 0, 0, 0, nil, true},
		{2024, 1, 31, 24, 0, 0, nil, true},
		{2024, 1, 31, 23, 60, 0, nil, true},
		{2024, 1, 31, 23, 59, 60, nil, true},
	}

	for _, tt := range tests {
		t.Run("TestNewDate", func(t *testing.T) {
			d, err := datetime.New(tt.year, tt.month, tt.day, tt.hour, tt.minute, tt.second)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, d)
			}
		})
	}

	t.Run("Test Now", func(t *testing.T) {
		d := datetime.Now()
		assert.NotNil(t, d)
	})
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		date     *datetime.DateTime
		expected bool
	}{
		{&datetime.DateTime{Year: 2024, Month: 12, Day: 21, DateTime: time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC)}, true},  // Saturday
		{&datetime.DateTime{Year: 2024, Month: 12, Day: 22, DateTime: time.Date(2024, 12, 22, 0, 0, 0, 0, time.UTC)}, true},  // Sunday
		{&datetime.DateTime{Year: 2024, Month: 12, Day: 23, DateTime: time.Date(2024, 12, 23, 0, 0, 0, 0, time.UTC)}, false}, // Monday
	}

	for _, tt := range tests {
		t.Run("TestIsWeekend", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date.IsWeekend())
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		date1    *datetime.DateTime
		date2    *datetime.DateTime
		expected int
	}{
		{&datetime.DateTime{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)},
			&datetime.DateTime{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, 0},
		{&datetime.DateTime{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)},
			&datetime.DateTime{Year: 2024, Month: 12, Day: 26, DateTime: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC)}, -1}, // 25th < 26th
		{&datetime.DateTime{Year: 2024, Month: 12, Day: 26, DateTime: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC)},
			&datetime.DateTime{Year: 2024, Month: 12, Day: 25, DateTime: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)}, 1}, // 26th > 25th
	}

	for _, tt := range tests {
		t.Run("TestCompare", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date1.Compare(tt.date2))
		})
	}
}

func TestDateTimeEqual(t *testing.T) {
	tests := []struct {
		date1    *datetime.DateTime
		date2    *datetime.DateTime
		expected bool
	}{
		// Test 1: Equal DateTime instances
		{
			date1:    &datetime.DateTime{Year: 2024, Month: 3, Day: 31},
			date2:    &datetime.DateTime{Year: 2024, Month: 3, Day: 31},
			expected: true,
		},
		// Test 2: Different DateTime instances (different day)
		{
			date1:    &datetime.DateTime{Year: 2024, Month: 3, Day: 31},
			date2:    &datetime.DateTime{Year: 2024, Month: 3, Day: 30},
			expected: false,
		},
		// Test 3: Different DateTime instances (different month)
		{
			date1:    &datetime.DateTime{Year: 2024, Month: 3, Day: 31},
			date2:    &datetime.DateTime{Year: 2024, Month: 4, Day: 1},
			expected: false,
		},
		// Test 4: Different DateTime instances (different year)
		{
			date1:    &datetime.DateTime{Year: 2024, Month: 3, Day: 31},
			date2:    &datetime.DateTime{Year: 2025, Month: 3, Day: 31},
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

func TestBetween(t *testing.T) {
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

func TestDateTimeFromString(t *testing.T) {
	validDate, _ := datetime.New(2025, 1, 31, 23, 27, 15)
	tests := []struct {
		date         string
		expectedErr  bool
		expectedDate datetime.Interface
	}{
		{"2025-01-31", true, nil},
		{"2025-01-31 23:27:15", false, validDate},
		{"2025-01-31 24:27:15", true, nil},
		{"2025-01-31 23:60:15", true, nil},
		{"2025-01-31 23:27:65", true, nil},
		{"2025-02-31", true, nil},
		{"2025-13-32", true, nil},
		{"-2025-06-31", true, nil},
		{"invalid", true, nil},
	}

	for _, tt := range tests {
		t.Run("TestDaysInMonthByDate: "+tt.date, func(t *testing.T) {
			d, err := datetime.FromString(tt.date)
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

func TestRange(t *testing.T) {
	t.Run("New Range", func(t *testing.T) {
		tests := []struct {
			rangeStr         string
			testDate         any
			expectedErr      bool
			expectedValidErr bool
			expected         func() any
		}{
			{"[2025-01-31, 2026-01-31]", "2025-05-06", true, false, func() any {
				val, err := datetime.NewRangeStrict(
					"2025-01-31", "2026-01-31")
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", false, false, func() any {
				val, err := datetime.NewRange(
					"2025-01-31 14:15:16", "2026-01-31 17:18:19", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", false, false, func() any {
				val, err := datetime.NewRange(
					"", "2026-01-31 17:18:19", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, ]", "2025-05-06 20:21:22", false, false, func() any {
				val, err := datetime.NewRange(
					"2025-01-31 14:15:16", "", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, ]", time.Date(2025, 5, 6, 20, 21, 22, 0, time.UTC), false, false, func() any {
				val, err := datetime.NewRange(
					"2025-01-31 14:15:16", "", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31, 2026-01-31]", "2025-05-06", true, false, nil},
			{"[2025-01-31 14:15:16, 2026-01-31]", "2025-05-06", true, false, nil},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", 2025, false, true, func() any {
				val, err := datetime.NewRange(
					"2025-01-31 14:15:16", "2026-01-31 17:18:19", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[, ]", datetime.Now(), false, true, func() any {
				val, err := datetime.NewRange(
					"", "", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.NoError(t, err)
				return val
			}},
			{"[2025-01-31 14:15:16, 2026-02-31 14:15:16]", datetime.Now(), true, true, func() any {
				val, err := datetime.NewRange(
					"2025-01-31 14:15:16", "2026-02-31 14:15:16", datetime.RangeStart("["), datetime.RangeEnd("]"))
				assert.Error(t, err)
				return val
			}},
		}

		for _, tt := range tests {
			t.Run("New Range: "+tt.rangeStr+fmt.Sprintf(", valid date: %v", tt.testDate), func(t *testing.T) {
				if tt.expectedErr {
					r, err := datetime.RangeFromString(tt.rangeStr)
					assert.Error(t, err)
					assert.Nil(t, r)

				} else {
					r, err := datetime.RangeFromString(tt.rangeStr)
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
			{"[2025-01-31, 2026-01-31]", "2025-05-06", true, false},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", false, true},
			{"[, 2026-01-31 17:18:19]", "2025-05-06 20:21:22", false, true},
			{"[2025-01-31 14:15:16, ]", "2025-05-06 20:21:22", false, true},
			{"[2025-01-31 14:15:16, ]", time.Date(2025, 5, 6, 20, 21, 22, 0, time.UTC), false, true},
			{"[2025-01-31, 2026-01-31]", "2025-05-06", true, false},
			{"[2025-01-31 14:15:16, 2026-01-31]", "2025-05-06", true, false},
			{"[2025-01-31 14:15:16, 2026-01-31 17:18:19]", 2025, false, false},
			{"[, ]", datetime.Now(), false, false},
			{"[2, ]", datetime.Now(), true, false},
		}

		for _, tt := range tests {
			t.Run("New Range: "+tt.rangeStr+fmt.Sprintf(", valid date: %v", tt.testDate), func(t *testing.T) {
				if tt.expectedErr {
					r, err := datetime.RangeFromString(tt.rangeStr)
					assert.Error(t, err)
					assert.Nil(t, r)

				} else {
					r, err := datetime.RangeFromString(tt.rangeStr)
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
			{true, "[", "2025-01-31", "2026-01-31", "]", nil},
			{false, "[", "2025-01-31 14:15:16", "2026-01-31 17:18:19", "]", func() any {
				val, err := datetime.NewRangeStrict("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "", "2026-01-31 17:18:19", "]", func() any {
				val, err := datetime.NewRangeStrict("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "", "]", func() any {
				val, err := datetime.NewRangeStrict("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "2026-01-31 17:18:19", "]", func() any {
				val, err := datetime.NewRangeStartOptional("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "", "2026-01-31 17:18:19", "]", func() any {
				val, err := datetime.NewRangeStartOptional("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "", "]", func() any {
				val, err := datetime.NewRangeStartOptional("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "2026-01-31 17:18:19", ")", func() any {
				val, err := datetime.NewRangeOptional("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "", "2026-01-31 17:18:19", ")", func() any {
				val, err := datetime.NewRangeOptional("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "(", "2025-01-31 14:15:16", "", ")", func() any {
				val, err := datetime.NewRangeOptional("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "2026-01-31 17:18:19", ")", func() any {
				val, err := datetime.NewRangeStartStrict("2025-01-31 14:15:16", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "", "2026-01-31 17:18:19", ")", func() any {
				val, err := datetime.NewRangeStartStrict("", "2026-01-31 17:18:19")
				assert.NoError(t, err)
				return val
			}},
			{false, "[", "2025-01-31 14:15:16", "", ")", func() any {
				val, err := datetime.NewRangeStartStrict("2025-01-31 14:15:16", "")
				assert.NoError(t, err)
				return val
			}},
		}

		for _, tt := range tests {
			t.Run("New Range brackets: "+fmt.Sprintf("From: %s%s, To: %s%s", tt.start, tt.from, tt.to, tt.end), func(t *testing.T) {
				r, err := datetime.NewRange(tt.from, tt.to, datetime.RangeStart(tt.start), datetime.RangeEnd(tt.end))

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
