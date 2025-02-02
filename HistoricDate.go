package datetime

type Historic string

var (
	BeforeChrist Historic = "bc"
	AfterChrist  Historic = "ac"
)

type HistoricDate struct {
	historic Historic
	date     DateTime
}
