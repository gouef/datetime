package datetime

import "time"

type Interface interface {
	ToString() string
	FromString(value string) (Interface, error)
	Time() time.Time
	Equal(u Interface) bool
	Between(start, end Interface) bool
	Before(u Interface) bool
	After(u Interface) bool
	Compare(u Interface) int
}
