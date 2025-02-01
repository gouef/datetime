package datetime

type RangeStart string
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

type Range struct {
	start RangeStart
	end   RangeEnd
}

func NewRange(start RangeStart, end RangeEnd) *Range {
	return &Range{
		start: start,
		end:   end,
	}
}

func NewRangeOptional() *Range {
	return NewRange(RangeStartOptional, RangeEndOptional)
}

func NewRangeStrict() *Range {
	return NewRange(RangeStartStrict, RangeEndStrict)
}

func NewRangeStartStrict() *Range {
	return NewRange(RangeStartStrict, RangeEndOptional)
}

func NewRangeStartOptional() *Range {
	return NewRange(RangeStartOptional, RangeEndStrict)
}
