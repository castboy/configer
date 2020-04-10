package structure

import (
	"fmt"
	"github.com/shopspring/decimal"
)

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
	QuoteOnly: "quote only",
	FullTrade: "full trade",
	NotSupport: "not support",
}

func (st SymbolStatus) String() string {
	return fmt.Sprintf("symbol trade right is `%s`", SymbolStatusMsg[st])
}

func (sb *Symbol) FormatCheck() error {
	return nil
}

func (sb *Symbol) AutoCondition() (cond string) {
	return
}