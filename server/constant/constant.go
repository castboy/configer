package constant

// MarginCalcMode Group MarginCalcMode
type MarginCalcMode int32

const (
	DoubleLegs MarginCalcMode = iota
	LargerLeg
	NetLeg
)

// GroupTradeType group trade type
type GroupTradeType int32

const (
	TestBook GroupTradeType = iota
	ABook
	BBook
)
