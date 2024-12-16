package datetime

import (
	"time"
)

func GetEaster(year int) time.Time {
	easter := Calculate(year)
	return easter
}

func GetMonday(year int) time.Time {
	easter := GetEaster(year)
	return easter.AddDate(0, 0, 1)
}

func GetGoodFriday(year int) time.Time {
	easter := GetEaster(year)
	return easter.AddDate(0, 0, -2)
}

func Calculate(year int) time.Time {
	s1, s2, d, e, a := calculatableVars(year)

	var easterSunday time.Time
	switch {
	case s1 >= 22 && s1 <= 31:
		easterSunday = GetDate(year, 3, s1)
	case s2 == 25 && d == 28 && e == 6 && a > 10:
		easterSunday = GetDate(year, 4, 18)
	case s2 <= 25:
		easterSunday = GetDate(year, 4, s2)
	default:
		easterSunday = GetDate(year, 4, s2-7)
	}

	return easterSunday
}

func calculatableVars(year int) (int, int, int, int, int) {
	a, b, c := cyclesVar(year)
	m, n := easterVar(year)
	d := ((19 * a) + m) % 30
	e := (n + (2 * b) + (4 * c) + (6 * d)) % 7

	s1 := 22 + d + e
	s2 := d + e - 9

	return s1, s2, d, e, a
}

func cyclesVar(year int) (int, int, int) {
	return year % 19, year % 4, year % 7
}

func easterVar(year int) (int, int) {
	if between(year, 1700, 1799) {
		return 23, 3
	} else if between(year, 1800, 1899) {
		return 23, 4
	} else if between(year, 1900, 2099) {
		return 24, 5
	}

	return 1, 1
}

func between(number int, low int, up int) bool {
	return number >= low && number <= up
}
