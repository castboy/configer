package structure

import "fmt"

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

// map holiday table.
type Holiday struct {
	ID          int             `json:"id" xorm:"id"`
	Enable      bool            `json:"enable" xorm:"enable"`
	Date        string          `json:"date" xorm:"date"`
	From        string          `json:"from" xorm:"from"`
	To          string          `json:"to" xorm:"to"`
	Category    HolidayCategory `json:"category" xorm:"category"`
	Symbol      string          `json:"symbol" xorm:"symbol"`
	SymbolID    int             `xorm:"-"`
	Description string          `json:"description" xorm:"description"`
}

type HolidayCategory int

const (
	HolidayAll HolidayCategory = iota
	HolidaySecurity
	HolidaySymbol
	HolidaySource
)

func (ho *Holiday) FormatCheck() error {
	return nil
}

func (ho *Holiday) IndexCheck() error {
	return nil
}

func (ho *Holiday) AutoCondition() (cond string) {
	return fmt.Sprintf("id = %d", ho.ID)
}


