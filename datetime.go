package datetime

import (
	"github.com/Gouef/utils"
	"time"
)

type Date struct {
	Year     int
	Month    int `validate:"min=1,max=12"`
	Day      int `validate:"min=1,max=31"`
	DateTime time.Time
}

func NewDate(year int, month int, day int) *Date {
	return &Date{Year: year, Month: month, Day: day, DateTime: GetDate(year, month, day)}
}

func GetDate(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (d *Date) IsWeekend() bool {
	weekendDays := []time.Weekday{time.Sunday, time.Saturday}
	return utils.InArray(d.DateTime.Weekday(), weekendDays)
}

func (d *Date) Time() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

// Compare compares the date instant d with u. If d is before u, it returns -1;
// if d is after u, it returns +1; if they're the same, it returns 0.
func (d Date) Compare(u Date) int {
	return d.Time().Compare(u.Time())
}

func (d Date) Equal(u Date) bool {
	return d.Time().Equal(u.Time())
}

func ListDaysInMonth(year int, month int) []int {
	days := make([]int, DaysInMonth(year, month))

	for i := range days {
		days[i] = i + 1
	}
	return days
}

func DaysInMonth(year int, month int) int {
	return DaysInMonthByDate(GetDate(year, month, 1))
}

func DaysInMonthByDate(t time.Time) int {
	y, m, _ := t.Date()
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
