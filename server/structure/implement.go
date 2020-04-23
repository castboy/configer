package structure

import (
	"fmt"
	"github.com/juju/errors"
	"time"
)

type MarketDST struct {
	MarketOwnerType MarketType `json:"market_owner_type" xorm:"market_type"`
	DST             DSTType    `xorm:"dst_type"`
}

func (md *MarketDST) FormatCheck() error {
	if md == nil {
		return errors.NotValidf("marketDST info is null")
	}

	if md.MarketOwnerType != NewYork && md.MarketOwnerType != London && md.MarketOwnerType != Sydney && md.MarketOwnerType != HK {
		return errors.NotValidf("marketOwnerType: %d", md.MarketOwnerType)
	}

	if md.DST != DST && md.DST != DSTNone && md.DST != None {
		return errors.NotValidf("dstType: %d", md.DST)
	}

	return nil
}
func (md *MarketDST) IndexCheck() error {
	return nil
}

func (md *MarketDST) NotFoundError() error {
	return errors.NotFoundf("MarketDST, MarketOwnerType: %d", md.MarketOwnerType)
}

func (md *MarketDST) ExportCondition() (cond string) {
	return "1"
}

func (md *MarketDST) UpdateCondition() (cond string) {
	return fmt.Sprintf("`market_type` = %d", md.MarketOwnerType)
}

func (md *MarketDST) DeleteCondition() (cond string) {
	return md.UpdateCondition()
}

func (md *MarketDST) GetCondition() (cond string) {
	return md.UpdateCondition()
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
	if fs.Sl.Symbol == "" {
		return errors.NotValidf("Symbol")
	}

	return nil
}

func (fs *FullSymbolName) NotFoundError() error {
	return errors.NotFoundf("FullSymbolName, Symbol: %s, Leverage: %d", fs.Sl.Symbol, fs.Sl.Leverage)
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
	if ho == nil {
		return errors.NotValidf("holiday info is null")
	}

	_, err := time.ParseInLocation("2006-01-02", ho.Date, time.UTC)
	if err != nil {
		return errors.NotValidf("date, %v", ho.Date)
	}

	_, err = time.ParseInLocation("15:04:05", ho.From, time.UTC)
	if err != nil {
		return errors.NotValidf("from, %v", ho.From)
	}

	_, err = time.ParseInLocation("15:04:05", ho.To, time.UTC)
	if err != nil {
		return errors.NotValidf("to, %v", ho.To)
	}

	// TODO

	return nil
}

func (ho *Holiday) IndexCheck() error {
	if ho.Date == "" {
		return errors.NotValidf("Date")
	}

	return nil
}

func (ho *Holiday) NotFoundError() error {
	return nil
}

func (ho *Holiday) ExportCondition() (cond string) {
	return "1"
}

func (ho *Holiday) UpdateCondition() (cond string) {
	return fmt.Sprintf("id = %d", ho.ID)
}

func (ho *Holiday) DeleteCondition() (cond string) {
	return ho.UpdateCondition()
}

func (ho *Holiday) GetCondition() (cond string) {
	return ho.UpdateCondition()
}
