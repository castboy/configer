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
	return se.SourceID
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
