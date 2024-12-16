package tests

import (
	"github.com/gouef/datetime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetEaster(t *testing.T) {
	tests := []struct {
		year     int
		expected time.Time
		isError  bool
	}{
		{2024, time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), false}, // 2024 Easter Sunday
		{2023, time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC), false},  // 2023 Easter Sunday
		{2022, time.Date(2022, 4, 17, 0, 0, 0, 0, time.UTC), false}, // 2022 Easter Sunday
		{2021, time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC), false},  // 2021 Easter Sunday
		{1801, time.Date(1801, 4, 5, 0, 0, 0, 0, time.UTC), false},  // 2021 Easter Sunday
		{1000, time.Date(1000, 4, 16, 0, 0, 0, 0, time.UTC), true},  // Test for invalid year (example year where calculation fails)
	}

	for _, tt := range tests {
		t.Run("TestGetEaster", func(t *testing.T) {
			easter := datetime.GetEaster(tt.year)
			if tt.isError {
				assert.Equal(t, tt.expected, easter)
			} else {
				assert.Equal(t, tt.expected, easter)
			}
		})
	}
}

func TestGetMonday(t *testing.T) {
	tests := []struct {
		year     int
		expected time.Time
	}{
		{2024, time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)},  // 2024 Easter Monday
		{2023, time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)}, // 2023 Easter Monday
		{2022, time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC)}, // 2022 Easter Monday
		{2021, time.Date(2021, 4, 5, 0, 0, 0, 0, time.UTC)},  // 2021 Easter Monday
	}

	for _, tt := range tests {
		t.Run("TestGetMonday", func(t *testing.T) {
			monday := datetime.GetMonday(tt.year)
			assert.Equal(t, tt.expected, monday)
		})
	}
}

func TestGetGoodFriday(t *testing.T) {
	tests := []struct {
		year     int
		expected time.Time
	}{
		{2024, time.Date(2024, 3, 29, 0, 0, 0, 0, time.UTC)}, // 2024 Good Friday
		{2023, time.Date(2023, 4, 7, 0, 0, 0, 0, time.UTC)},  // 2023 Good Friday
		{2022, time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC)}, // 2022 Good Friday
		{2021, time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)},  // 2021 Good Friday
	}

	for _, tt := range tests {
		t.Run("TestGetGoodFriday", func(t *testing.T) {
			goodFriday := datetime.GetGoodFriday(tt.year)
			assert.Equal(t, tt.expected, goodFriday)
		})
	}
}
