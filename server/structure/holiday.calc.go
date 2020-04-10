package structure

type HolidayCalc struct {
	ID int
	DateSymbol
	TimeSpans []*TimeSpan
}

type DateSymbol struct {
	Date,
	Symbol string
}

type TimeSpan struct {
	From,
	To string
}

func (hc *HolidayCalc) FormatCheck() error {
	return nil
}

func (hc *HolidayCalc) AutoCondition() (cond string) {
	return
}