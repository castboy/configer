package structure

import "fmt"

type IDor interface {
	GetID() int
}

// map session table.
type Session struct {
	ID       int                `xorm:"id autoincr"`
	SourceID int                `xorm:"source_id"`
	Type     SessionType        `xorm:"type"`
	Dst      DSTType            `xorm:"dst_type"`
	Session  map[int32][]string `xorm:"session"`
}

type SessionType int

const (
	Quote SessionType = iota
	Trade
	SessionTypeLength
)

type DSTType int

const (
	DST DSTType = iota
	DSTNone
	None
	DSTTypeLength
)

func (se *Session) FormatCheck() error {
	return nil
}

func (se *Session) IndexCheck() error {
	return nil
}

func (se *Session) AutoCondition() (cond string) {
	return fmt.Sprintf("`type` = %d and `dst_type` = %d", se.Type, se.Dst)
}

func (se *Session) GetID() int {
	return se.ID
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

func (ho *Holiday) GetID() int {
	return ho.ID
}

//
type ConvSymbol struct {
	*ConvInfo
	ConvType
	SourceID int
}

type ConvInfo struct {
	ConvNeed,
	ConvMultiply bool
	ConvSymbol string
}

type ConvType int

const (
	MarginConv ConvType = iota
	ProfitConv
	ConvTypeLength
)

func (cs *ConvSymbol) FormatCheck() error {
	return nil
}

func (cs *ConvSymbol) IndexCheck() error {
	return nil
}

func (cs *ConvSymbol) AutoCondition() (cond string) {
	return
}

func (cs *ConvSymbol) GetID() int {
	return cs.SourceID
}
