package datetime

type RangeStart string
type RangeEnd string

var (
	RANGE_START_STRICT   RangeStart = "["
	RANGE_START_OPTIONAL RangeStart = "("
	RANGE_END_STRICT     RangeEnd   = "]"
	RANGE_END_OPTIONAL   RangeEnd   = ")"
)