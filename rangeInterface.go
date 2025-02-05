package datetime

// RangeStart start bracket
type RangeStart string

// RangeEnd end bracket
type RangeEnd string

var (
	// RangeStartStrict can not be equal
	RangeStartStrict RangeStart = "["
	// RangeStartOptional can be equal
	RangeStartOptional RangeStart = "("
	// RangeEndStrict can not be equal
	RangeEndStrict RangeEnd = "]"
	// RangeEndOptional can be equal
	RangeEndOptional RangeEnd = ")"
)

const (
	YearRegexp     = `(\d+)`
	MonthRegexp    = `(0[1-9]|1[0-2])`
	DayRegexp      = `(0[1-9]|[12][0-9]|3[01])`
	HourRegexp     = `(0[0-9]|1[0-9]|2[0-3])`
	MinuteRegexp   = `[0-5][0-9]`
	SecondRegexp   = `[0-5][0-9]`
	DateRegexp     = YearRegexp + `-` + MonthRegexp + `-` + DayRegexp
	TimeRegexp     = `(` + HourRegexp + `):(` + MinuteRegexp + `):(` + SecondRegexp + `)`
	DateTimeRegexp = `((` + DateRegexp + `) (` + TimeRegexp + `))`
	RangeRegexp    = `^([\[\(])` + DateTimeRegexp + `?\s*,\s*` + DateTimeRegexp + `?([\]\)])$`
)

type RangeInterface interface {
	Start() RangeStart
	End() RangeEnd
	From() ValueInterface
	To() ValueInterface
	String() string
	Is(value any) bool
}
