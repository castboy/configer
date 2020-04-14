package structure

//implement `Checkor` ,`Xormor` interfaces.

type MarketDST struct {
	MarketOwnerType MarketType `json:"market_owner_type" xorm:"market_type"`
	DST             DSTType    `xorm:"dst_type"`
}

func (md *MarketDST) FormatCheck() error {
	return nil
}
func (md *MarketDST) IndexCheck() error {
	return nil
}

func (md *MarketDST) AutoCondition() (cond string) {
	return
}

// holiday.calc
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

func (hc *HolidayCalc) IndexCheck() error {
	return nil
}

func (hc *HolidayCalc) AutoCondition() (cond string) {
	return
}

// full.symbol.name
type FullSymbolName struct {
	Sl       SymbolLeverage
	FullName string
}

type SymbolLeverage struct {
	Symbol   string
	Leverage int32
}

func (fs *FullSymbolName) FormatCheck() error {
	return nil
}

func (fs *FullSymbolName) IndexCheck() error {
	return nil
}

func (fs *FullSymbolName) AutoCondition() (cond string) {
	return
}


