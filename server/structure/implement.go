package structure

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/shopspring/decimal"
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
	return errors.NotFoundf("Date: %v", ho.Date)
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

//congroup security
//ConGroupSec
type ConGroupSec struct {
	ID             int             `xorm:"'id' pk " `
	GroupId        int             `xorm:"group_id"`
	SecurityId     int             `xorm:"security_id"`
	EnableSecurity bool            `xorm:"enable_security"`
	EnableTrade    bool            `xorm:"enable_trade"`
	LotMin         decimal.Decimal `xorm:"lot_min"`
	LotMax         decimal.Decimal `xorm:"lot_max"`
	LotStep        decimal.Decimal `xorm:"lot_step"`
	SpreadDiff     int             `xorm:"spread_diff"`
	Commission     decimal.Decimal `xorm:"commission"`
}

func (cgs *ConGroupSec) FormatCheck() error {
	if cgs == nil {
		return errors.NotValidf("session info is null")
	}

	if cgs.GroupId <= 0 {
		return errors.NotValidf("groupID, %v", cgs.GroupId)
	}

	if cgs.SecurityId <= 0 {
		return errors.NotValidf("SecurityId, %v", cgs.SecurityId)
	}

	if cgs.LotMin.LessThan(decimal.Zero) {
		return errors.NotValidf("LotMin, %v", cgs.LotMin)
	}

	if cgs.LotMax.LessThan(decimal.Zero) {
		return errors.NotValidf("LotMax, %v", cgs.LotMax)
	}

	if cgs.LotStep.LessThan(decimal.Zero) {
		return errors.NotValidf("LotStep, %v", cgs.LotStep)
	}

	return nil
}

func (cgs *ConGroupSec) IndexCheck() error {
	if cgs.GroupId == 0 {
		return errors.NotValidf("GroupID, %v", cgs.GroupId)
	}

	if cgs.SecurityId == 0 {
		return errors.NotValidf("SecurityID, %v", cgs.SecurityId)
	}

	return nil
}

func (cgs *ConGroupSec) NotFoundError() error {
	return errors.NotFoundf("ConGroupSec, groupID: %d, securityID:%d", cgs.GroupId, cgs.SecurityId)
}

func (cgs *ConGroupSec) ExportCondition() (cond string) {
	return "1"
}

func (cgs *ConGroupSec) UpdateCondition() (cond string) {
	if cgs.ID != 0 {
		return fmt.Sprintf("`id` = %d", cgs.ID)
	}

	return fmt.Sprintf("`group_id` = %d and `security_id` = %d", cgs.GroupId, cgs.SecurityId)
}

func (cgs *ConGroupSec) DeleteCondition() (cond string) {
	return cgs.UpdateCondition()
}

func (cgs *ConGroupSec) GetCondition() (cond string) {
	return cgs.UpdateCondition()
}
