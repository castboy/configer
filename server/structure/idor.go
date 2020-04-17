package structure

import (
	"configer/server/constant"
	"fmt"
	"github.com/juju/errors"
	"regexp"
	"sort"
	"strings"
)

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
	if se == nil {
		return constant.NewErr(constant.ArgsErr, errors.NotValidf("session info"))
	}

	if se.Type != Quote && se.Type != Trade {
		return constant.NewErr(constant.TradeErr, errors.NotValidf("type, %v", se.Type))
	}

	if se.Dst != None && se.Dst != DST && se.Dst != DSTNone {
		return constant.NewErr(constant.ArgsErr, errors.NotValidf("dstType, %v", se.Dst))
	}

	if se.Session == nil || len(se.Session) == 0 {
		return constant.NewErr(constant.ArgsErr, errors.NotValidf("session, %v", se.Session))
	}

	// sure "00:00-00:00" format
	for weekday, _ := range se.Session {
		for i, _ := range se.Session[weekday] {
			matched, _ := regexp.MatchString(`^(20|21|22|23|24|[0-1]\d):[0-5]\d-(20|21|22|23|24|[0-1]\d):[0-5]\d$`, se.Session[weekday][i])
			if !matched {
				return constant.NewErr(constant.ArgsErr, errors.NotValidf("session format: %s", se.Session[weekday][i]))
			}
		}
	}

	// "01:00-00:00" -> "01:00->24:00"
	for weekday, _ := range se.Session {
		for i, _ := range se.Session[weekday] {
			ts := strings.Split(se.Session[weekday][i], "-")
			if ts[0] != "00:00" && ts[1] == "00:00" {
				se.Session[weekday][i] = ts[0] + "-" + "24:00"
			}
		}
	}

	// sure no overlap
	for weekday, _ := range se.Session {
		sort.Strings(se.Session[weekday])

		s := []string{}
		for i := range se.Session[weekday] {
			if se.Session[weekday][i] != "00:00-00:00" {
				s = append(s, strings.Split(se.Session[weekday][i], "-")...)
			}
		}

		end := len(s)

		for i := 0; i < end-1; i++ {
			if s[i] >= s[i+1] {
				return errors.NotValidf("session format: %v, sessionType: %d", se.Session[weekday], se.Type)
			}
		}
	}

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
