package main

import (
	"configer/server/base"
	"configer/server/structure"
	"configer/server/utils"
	"github.com/juju/errors"
	"strings"
	"time"
)

func ExportSessions(sourceName string, dstType structure.DSTType, sessionType structure.SessionType) (map[int32][]string, error) {
	id, err := GetSourceIDByName(sourceName)
	if err != nil {
		return nil, err
	}

	i, err := base.Export(base.NewSessioner(&structure.Session{Type: sessionType, Dst: dstType}))
	if err != nil {
		return nil, nil
	}

	ss := i.(map[int]structure.IDor)
	for j := range ss {
		s := ss[j].(*structure.Session)
		if s.SourceID == id {
			s.Session = utils.OrderAndFill(s.Session)
			return s.Session, nil
		}
	}

	return nil, nil
}

// check below.
func SetSession(sess *structure.Session) (err error) {
	_, err = base.Update(base.NewSessioner(sess))
	return
}

func GetHolidays() (hs []*structure.Holiday, err error) {
	i, err := base.Export(base.NewHolidayer(&structure.Holiday{}))
	if err != nil {
		return
	}

	hos := i.(map[string]map[int]*structure.Holiday)
	for date := range hos {
		for id := range hos[date] {
			hs = append(hs, hos[date][id])
		}
	}

	return
}

func GetHolidayByID(ID int) (ho *structure.Holiday, err error) {
	i, err := base.Get(base.NewHolidayer(&structure.Holiday{ID: ID}))
	if err != nil {
		return
	}

	return i.(*structure.Holiday), nil
}

func InsertHoliday(holi *structure.Holiday) error {
	return base.Insert(base.NewHolidayer(holi))
}

func UpdateHolidayByID(ID int, holi *structure.Holiday) error {
	holi.ID = ID
	_, err := base.Update(base.NewHolidayer(holi))
	return err
}

func DeleteHolidayByID(ID int) error {
	_, err := base.Delete(base.NewHolidayer(&structure.Holiday{ID: ID}))
	return err
}


func IsQuotable(symb *structure.Symbol) bool {
	if !holidayCanTrade(symb) {
		return false
	}

	return sessionCanQuoteTrade(symb, structure.Quote)
}

func IsTradable(symb *structure.Symbol) bool {
	if !holidayCanTrade(symb) {
		return false
	}

	return sessionCanQuoteTrade(symb, structure.Trade)
}


func appendSymbolID(ho *structure.Holiday) error {
	switch ho.Category {
	case structure.HolidayAll:

	case structure.HolidaySecurity:
		i, err := base.Get(base.NewSecurityer(&structure.Security{SecurityName: ho.Symbol}))
		if err != nil {
			return err
		}

		se := i.(*structure.Security)
		ho.SymbolID = se.ID

	case structure.HolidaySource:
		i, err := base.Get(base.NewSourcer(&structure.Source{Source: ho.Symbol}))
		if err != nil {
			return err
		}

		src := i.(*structure.Source)
		ho.SymbolID = src.ID

	case structure.HolidaySymbol:
		i, err := base.Get(base.NewSymboler(&structure.Symbol{Symbol: ho.Symbol}))
		if err != nil {
			return err
		}

		symb := i.(*structure.Symbol)
		ho.SymbolID = symb.ID

	default:
		return errors.NotValidf("holiday category: %d", ho.Category)
	}

	return nil
}

func holidayCanTrade(symb *structure.Symbol) bool {
	date := time.Now().UTC().Format("2006-01-02")
	holiday := &structure.Holiday{Date: date}
	i, err := base.Get(base.NewHolidayer(holiday))
	if err != nil {
		return true
	}

	now := time.Now().UTC().Format("15:04:05")
	timeCanTrade := func(from, to string) bool {
		if now >= from && now <= to {
			return true
		}

		return false
	}

	hos := i.(map[int]*structure.Holiday)
	for j := range hos {
		if !hos[j].Enable {
			continue
		}

		switch hos[j].Category {
		case structure.HolidayAll:
			if !timeCanTrade(hos[j].From, hos[j].To) {
				return false
			}
		case structure.HolidaySecurity:
			if hos[j].SymbolID == symb.SecurityID {
				if !timeCanTrade(hos[j].From, hos[j].To) {
					return false
				}
			}
		case structure.HolidaySource:
			if hos[j].SymbolID == symb.SourceID {
				if !timeCanTrade(hos[j].From, hos[j].To) {
					return false
				}
			}
		case structure.HolidaySymbol:
			if hos[j].SymbolID == symb.ID {
				if !timeCanTrade(hos[j].From, hos[j].To) {
					return false
				}
			}
		}
	}

	return true
}

func sessionCanQuoteTrade(symb *structure.Symbol, t structure.SessionType) bool {
	// get session
	i, err := base.Get(base.NewSourcer(&structure.Source{ID: symb.SourceID}))
	if err != nil {

	}

	src := i.(*structure.Source)

	j, err := base.Get(base.NewMarketDSTer(&structure.MarketDST{MarketOwnerType: src.MarketOwnerType}))
	if err != nil {

	}

	md := j.(*structure.MarketDST)

	k, err := base.Get(base.NewSessioner(&structure.Session{SourceID: symb.SourceID, Type: t, Dst: md.DST}))
	if err != nil {

	}

	ses := k.(*structure.Session)

	// judge
	gmt := time.Now().UTC()
	nowStr := gmt.Format("15:04:05")
	weekday := gmt.Weekday()

	for _, session := range ses.Session[int32(weekday)] {
		beginEnd := strings.Split(session, "-")
		if beginEnd[0] <= nowStr && nowStr < beginEnd[1] {
			return true
		}
	}

	return false
}
