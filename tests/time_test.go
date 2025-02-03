package tests

import (
	"github.com/gouef/datetime"
	"github.com/gouef/datetime/time"
	"github.com/stretchr/testify/assert"
	"testing"
	goTime "time"
)

func TestNewTime(t *testing.T) {
	tests := []struct {
		hour     int
		minute   int
		second   int
		expected *time.Time
		err      bool
	}{
		{20, 12, 25, &time.Time{Hour: 20, Minute: 12, Second: 25, DateTime: goTime.Date(0, 1, 1, 20, 12, 25, 0, goTime.UTC)}, false},
		{24, 2, 30, nil, true},   // Invalid day for February
		{20, 60, 10, nil, true},  // Invalid month (13)
		{20, 4, 60, nil, true},   // April has only 30 days
		{-20, 12, 25, nil, true}, // April has only 30 days
	}

	for _, tt := range tests {
		t.Run("TestNewDate", func(t *testing.T) {
			d, err := time.New(tt.hour, tt.minute, tt.second)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, d)
			}
		})
	}
}

func TestTimeCompare(t *testing.T) {
	tests := []struct {
		date1    *time.Time
		date2    *time.Time
		expected int
	}{
		{&time.Time{Hour: 2024, Minute: 12, Second: 25, DateTime: goTime.Date(2024, 12, 25, 0, 0, 0, 0, goTime.UTC)},
			&time.Time{Hour: 2024, Minute: 12, Second: 25, DateTime: goTime.Date(2024, 12, 25, 0, 0, 0, 0, goTime.UTC)}, 0},
		{&time.Time{Hour: 2024, Minute: 12, Second: 25, DateTime: goTime.Date(2024, 12, 25, 0, 0, 0, 0, goTime.UTC)},
			&time.Time{Hour: 2024, Minute: 12, Second: 26, DateTime: goTime.Date(2024, 12, 26, 0, 0, 0, 0, goTime.UTC)}, -1}, // 25th < 26th
		{&time.Time{Hour: 2024, Minute: 12, Second: 26, DateTime: goTime.Date(2024, 12, 26, 0, 0, 0, 0, goTime.UTC)},
			&time.Time{Hour: 2024, Minute: 12, Second: 25, DateTime: goTime.Date(2024, 12, 25, 0, 0, 0, 0, goTime.UTC)}, 1}, // 26th > 25th
	}

	for _, tt := range tests {
		t.Run("TestCompare", func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.date1.Compare(tt.date2))
		})
	}
}

func TestTimeEqual(t *testing.T) {
	tests := []struct {
		date1    *time.Time
		date2    *time.Time
		expected bool
	}{
		// Test 1: Equal DateTime instances
		{
			date1:    &time.Time{Hour: 20, Minute: 3, Second: 31},
			date2:    &time.Time{Hour: 20, Minute: 3, Second: 31},
			expected: true,
		},
		// Test 2: Different DateTime instances (different day)
		{
			date1:    &time.Time{Hour: 20, Minute: 3, Second: 31},
			date2:    &time.Time{Hour: 2024, Minute: 3, Second: 30},
			expected: false,
		},
		// Test 3: Different DateTime instances (different month)
		{
			date1:    &time.Time{Hour: 20, Minute: 3, Second: 31},
			date2:    &time.Time{Hour: 2024, Minute: 4, Second: 1},
			expected: false,
		},
		// Test 4: Different DateTime instances (different year)
		{
			date1:    &time.Time{Hour: 20, Minute: 3, Second: 31},
			date2:    &time.Time{Hour: 2025, Minute: 3, Second: 31},
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

func TestTimeBetween(t *testing.T) {
	date1, _ := time.New(20, 25, 1)
	date2, _ := time.New(20, 24, 1)
	date3, _ := time.New(20, 26, 1)

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

func TestTimeFromString(t *testing.T) {
	validDate, _ := time.New(23, 27, 15)
	tests := []struct {
		date         string
		expectedErr  bool
		expectedDate datetime.Interface
	}{
		{"2025-01-31", true, nil},
		{"2025-01-31 23:27:15", false, validDate},
		{"2025-02-31", true, nil},
		{"2025-13-32", true, nil},
		{"-2025-06-31", true, nil},
		{"invalid", true, nil},
	}

	for _, tt := range tests {
		t.Run("TestTimeFromString: "+tt.date, func(t *testing.T) {
			d, err := time.FromString(tt.date)
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
