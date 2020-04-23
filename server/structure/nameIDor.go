package structure

import (
	"configer/server/constant"
	"fmt"
	"github.com/juju/errors"
	"github.com/shopspring/decimal"
	"time"
)

type NameIDor interface {
	GetName() string
	GetID() int
}

// map symbol table.
type Symbol struct {
	ID            int             `json:"id" xorm:"id autoincr"`
	Index         int             `json:"index" xorm:"index"`
	Symbol        string          `json:"symbol" xorm:"symbol"`
	SourceID      int             `json:"source_id" xorm:"source_id"`
	Leverage      int32           `json:"leverage" xorm:"-"`
	SecurityID    int             `json:"security_id" xorm:"security_id"`
	MarginInitial decimal.Decimal `json:"margin_initial" xorm:"margin_initial"`
	MarginDivider decimal.Decimal `json:"margin_divider" xorm:"margin_divider"`
	Percentage    decimal.Decimal `json:"percentage" xorm:"percentage"`
	Status        SymbolStatus    `json:"status" xorm:"status"`
}

type (
	SymbolStatus int
)

const (
	QuoteOnly SymbolStatus = iota
	FullTrade
	NotSupport
)

var SymbolStatusMsg = map[SymbolStatus]string{
	QuoteOnly:  "quote only",
	FullTrade:  "full trade",
	NotSupport: "not support",
}

func (st SymbolStatus) String() string {
	return fmt.Sprintf("symbol trade right is `%s`", SymbolStatusMsg[st])
}

func (sb *Symbol) FormatCheck() error {
	if sb.Symbol == "" {
		return errors.NotValidf("Symbol, %v", sb.Symbol)
	}

	if sb.SourceID == 0 {
		return errors.NotValidf("SourceID, %v", sb.SourceID)
	}

	if sb.SecurityID == 0 {
		return errors.NotValidf("SecurityID, %v", sb.SecurityID)
	}

	if sb.Leverage == 0 {
		return errors.NotValidf("Leverage, %v", sb.Leverage)
	}

	if sb.MarginDivider.IsZero() {
		return errors.NotValidf("MarginDivider, %v", 0)
	}

	if sb.Percentage.IsZero() {
		return errors.NotValidf("Percentage, %v", 0)
	}

	return nil
}

func (sb *Symbol) IndexCheck() error {
	if sb.GetID() == 0 && sb.GetName() == "" {
		return errors.NotValidf("ID and name")
	}

	return nil
}

func (sb *Symbol) NotFoundError() error {
	return errors.NotFoundf("Symbol, SymbolName: %s, ID: %d", sb.Symbol, sb.ID)
}

func (sb *Symbol) ExportCondition() (cond string) {
	return "1"
}

func (sb *Symbol) UpdateCondition() (cond string) {
	if sb.Symbol != "" {
		cond = fmt.Sprintf("`symbol` = '%s'", sb.Symbol)
		return
	}

	if sb.ID != 0 {
		cond = fmt.Sprintf("where `id` = %d", sb.ID)
		return
	}

	return
}
func (sb *Symbol) DeleteCondition() (cond string) {
	return sb.UpdateCondition()
}
func (sb *Symbol) GetCondition() (cond string) {
	return sb.UpdateCondition()
}

func (sb *Symbol) GetName() string {
	return sb.Symbol
}

func (sb *Symbol) GetID() int {
	return sb.ID
}

// map source table.
type Source struct {
	ID              int             `json:"id" xorm:"id autoincr"`
	Source          string          `json:"source" xorm:"source"`
	SourceCN        string          `json:"source_cn" xorm:"source_cn"`
	SourceType      SourceType      `json:"source_type" xorm:"source_type"`
	Digits          int             `json:"digits" xorm:"digits"`
	Multiply        decimal.Decimal `json:"multiply" xorm:"multiply"`
	ContractSize    decimal.Decimal `json:"contract_size" xorm:"contract_size"`
	StopsLevel      int             `json:"stops_level" xorm:"stops_level"`
	ProfitMode      ProfitMode      `json:"profit_mode" xorm:"profit_mode"`
	Currency        string          `json:"currency" xorm:"currency"`
	ProfitCurrency  string          `json:"profit_currency" xorm:"profit_currency"`
	MarginMode      MarginMode      `json:"margin_mode" xorm:"margin_mode"`
	MarginCurrency  string          `json:"margin_currency" xorm:"margin_currency"`
	SwapType        SwapType        `json:"swap_type" xorm:"swap_type"`
	SwapLong        decimal.Decimal `json:"swap_long" xorm:"swap_long"`
	SwapShort       decimal.Decimal `json:"swap_short" xorm:"swap_short"`
	SwapCurrency    string          `josn:"swap_currency" xorm:"swap_currency"`
	Swap3Day        time.Weekday    `json:"swap_3_day" xorm:"swap_3_day"`
	MarketOwnerType MarketType      `json:"market_owner_type" xorm:"market_owner_type"`
}

type (
	ProfitMode int
	SwapType   int
	MarginMode int
	SourceType int
	MarketType int
)

const (
	ProfitForex ProfitMode = iota
	ProfitCfd
	ProfitFutures
)

const (
	ByPoints SwapType = iota
	ByMoney
	ByInterest
	ByMoneyInMarginCurrency
	ByInterestOfCfds
	ByInterestOfFutures
)

const (
	MarginForex MarginMode = iota
	MarginCfd
	MarginFutures
	MarginCfdIndex
	MarginCfdLeverage
)

const (
	SourceFx SourceType = iota
	SourceMetal
	SourceEnergy
	SourceIndex
	SourceCrypto
)

const (
	NewYork MarketType = iota
	London
	Sydney
	HK
)

func (src *Source) FormatCheck() error {
	if src == nil {
		return errors.NotValidf("source info is null")
	}

	if src.Source == "" {
		return errors.NotValidf("source, %v", src.Source)
	}

	if src.SourceType != SourceFx && src.SourceType != SourceMetal && src.SourceType != SourceEnergy &&
		src.SourceType != SourceIndex && src.SourceType != SourceCrypto {
		return errors.NotValidf("sourceType, %v", src.SourceType)
	}

	const DigitsMaxLimit = 5

	if src.Digits <= 0 || src.Digits > DigitsMaxLimit {
		return errors.NotValidf("digits, %v", src.Digits)
	}

	if src.ContractSize.Equal(decimal.NewFromFloat(0)) {
		return errors.NotValidf("contractSize, %v", src.ContractSize.String())
	}

	if src.StopsLevel == 0 {
		return errors.NotValidf("stopsLevel, %v", src.StopsLevel)
	}

	if src.ProfitMode != ProfitForex && src.ProfitMode != ProfitCfd {
		return errors.NotValidf("profitMode, %v", src.ProfitMode)
	}

	if src.Currency == "" {
		return errors.NotValidf("currency, %v", src.Currency)
	}

	if src.ProfitCurrency == "" {
		return errors.NotValidf("profitCurrency, %v", src.ProfitCurrency)
	}

	if src.MarginMode != MarginForex && src.MarginMode != MarginCfd && src.MarginMode != MarginFutures &&
		src.MarginMode != MarginCfdLeverage {
		return errors.NotValidf("marginMode")
	}

	if src.MarginCurrency == "" {
		return errors.NotValidf("marginCurrency")
	}

	if src.SwapType != ByPoints && src.SwapType != ByInterestOfCfds {
		return errors.NotValidf("swapType, %v", src.SwapType)
	}

	if src.SwapLong.IsZero() && src.SwapShort.IsZero() {
		return errors.NotValidf("swapLong, %v; swapShort, %v", src.SwapLong.String(), src.SwapShort.String())
	}

	return nil
}

func (src *Source) IndexCheck() error {
	if src.GetID() == 0 && src.GetName() == "" {
		return errors.NotValidf("ID and name")
	}

	return nil
}

func (src *Source) NotFoundError() error {
	return errors.NotFoundf("Source, SourceName: %s, ID: %d", src.Source, src.ID)
}

func (src *Source) ExportCondition() (cond string) {
	return "1"
}

func (src *Source) UpdateCondition() (cond string) {
	if src.Source != "" {
		cond = fmt.Sprintf("`source` = '%s'", src.Source)
		return
	}

	if src.ID != 0 {
		cond = fmt.Sprintf("where `id` = %d", src.ID)
		return
	}

	return
}
func (src *Source) DeleteCondition() (cond string) {
	return src.UpdateCondition()
}
func (src *Source) GetCondition() (cond string) {
	return src.UpdateCondition()
}

func (src *Source) GetName() string {
	return src.Source
}

func (src *Source) GetID() int {
	return src.ID
}

// map security table.
type Security struct {
	ID           int      `json:"id" xorm:"id"`
	SecurityName string   `json:"security_name" xorm:"security_name"`
	Description  string   `json:"description" xorm:"description"`
	Symbols      []string `json:"symbols" xorm:"-"`
}

func (sec *Security) FormatCheck() error {
	return nil
}

func (sec *Security) IndexCheck() error {
	if sec.GetID() == 0 && sec.GetName() == "" {
		return errors.NotValidf("ID and name")
	}

	return nil
}

func (sec *Security) NotFoundError() error {
	return errors.NotFoundf("Security, SecurityName: %s, ID: %d", sec.SecurityName, sec.ID)
}

func (sec *Security) ExportCondition() (cond string) {
	return "1"
}

func (sec *Security) UpdateCondition() (cond string) {
	if sec.SecurityName != "" {
		cond = fmt.Sprintf("`security_name` = '%s'", sec.SecurityName)
		return
	}

	if sec.ID != 0 {
		cond = fmt.Sprintf("where `id` = %d", sec.ID)
		return
	}

	return
}
func (sec *Security) DeleteCondition() (cond string) {
	return sec.UpdateCondition()
}
func (sec *Security) GetCondition() (cond string) {
	return sec.UpdateCondition()
}

func (sec *Security) GetName() string {
	return sec.SecurityName
}

func (sec *Security) GetID() int {
	return sec.ID
}

//Group
type AccountGroup struct {
	ID              int                     `xorm:"id autoincr"`
	Name            string                  `xorm:"name notnull"`
	DepositCurrency string                  `xorm:"deposit_currency notnull"`
	MarginStopOut   decimal.Decimal         `xorm:"margin_stop_out notnull"`
	MarginMode      constant.MarginCalcMode `xorm:"margin_mode notnull"`
	MarginCall      decimal.Decimal         `xorm:"margin_call"`
	TradeType       constant.GroupTradeType `xorm:trade_type notnull`
	IsChargeSwap    bool                    `xorm:"is_charge_swap"`
	Comment         string                  `xorm:"comment"`
}

func (grp *AccountGroup) FormatCheck() error {
	if grp == nil {
		return errors.NotValidf("AccountGroup info is null")
	}

	if grp.Name == "" {
		return errors.NotValidf("AccountGroup, %v", grp.Name)
	}

	if grp.DepositCurrency == "" {
		return errors.NotValidf("AccountGroup DepositCurrency is null")
	}

	if grp.MarginStopOut.LessThanOrEqual(decimal.Zero) {
		return errors.NotValidf("AccountGroup MarginStopOut is invalid")
	}

	if grp.MarginMode >= constant.MarginCalcModeEnd {
		return errors.NotValidf("AccountGroup MarginMode is invalid")
	}

	if grp.TradeType >= constant.GroupTradeTypeEnd {
		return errors.NotValidf("AccountGroup TradeType is invalid")
	}
	return nil
}

func (grp *AccountGroup) IndexCheck() error {
	if grp.GetID() == 0 && grp.GetName() == "" {
		return errors.NotValidf("ID and name")
	}

	return nil
}

func (grp *AccountGroup) NotFoundError() error {
	return errors.NotFoundf("AccountGroupName, AccountGroupID: %s, ID: %d", grp.Name, grp.ID)
}

func (sgrp *AccountGroup) ExportCondition() (cond string) {
	return "1"
}

func (grp *AccountGroup) UpdateCondition() (cond string) {
	if grp.Name != "" {
		cond = fmt.Sprintf("`name` = '%s'", grp.Name)
		return
	}

	if grp.ID != 0 {
		cond = fmt.Sprintf("`id` = %d", grp.ID)
		return
	}

	return
}
func (grp *AccountGroup) DeleteCondition() (cond string) {
	return grp.UpdateCondition()
}
func (grp *AccountGroup) GetCondition() (cond string) {
	return grp.UpdateCondition()
}

func (grp *AccountGroup) GetName() string {
	return grp.Name
}

func (grp *AccountGroup) GetID() int {
	return grp.ID
}
