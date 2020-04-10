package main

import (
	"configer/server/base"
	"configer/server/extend"
	"configer/server/structure"
	"fmt"
)

func main() {
	a := &structure.Symbol{Symbol: "AUDCAD"}
	exist, err := base.Get(base.NewSymboler(a))
	fmt.Println(a, exist, err)

	b := &structure.Symbol{Symbol: "AUDCAD"}
	num, err := base.Insert(base.NewSymboler(b))
	fmt.Println(num, err)
}

func Start() error {
	// cache symbol.
	symbol := &structure.Symbol{}
	symboler := base.NewSymboler(symbol)
	err := base.Cache(symboler)
	if err != nil {

	}

	// cache source.
	src := &structure.Source{}
	err = base.Cache(base.NewSourcer(src))
	if err != nil {

	}

	// cache session.
	sess := &structure.Session{}
	err = base.Cache(base.NewSessioner(sess))
	if err != nil {

	}

	// cache security.
	sec := &structure.Security{}
	err = base.Cache(base.NewSecurityer(sec))
	if err != nil {

	}

	// cache market_dst.
	md := &structure.MarketDST{}
	err = base.Cache(base.NewMarketDSTer(md))
	if err != nil {

	}

	// cache full_symbol_name.
	fsn := &structure.FullSymbolName{}
	err = extend.Cache(extend.NewFullSymbolNamer(symbol, fsn))
	if err != nil {

	}

	// cache conv_symbol
	cs := &structure.ConvSymbol{}
	err = extend.Cache(extend.NewConvSymboler(symbol, cs))
	if err != nil {

	}

	// cache holiday
	ho := &structure.Holiday{}
	holidayer := base.NewHolidayer(ho)
	err = base.Cache(holidayer)
	if err != nil {

	}

	// cache holiday.calc
	hs, err := base.Export(holidayer)
	if err != nil {

	}

	holidays := hs.([]structure.Holiday)
	for i := range holidays {
		hc := &structure.HolidayCalc{}
		hc.ID = holidays[i].ID
		hc.Date = holidays[i].Date
		hc.TimeSpans[0] = &structure.TimeSpan{From: holidays[i].From, To: holidays[i].To}

		switch holidays[i].Category {
		case structure.HolidayAll:
			err := extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{}, hc))
			if err != nil {

			}

		case structure.HolidaySecurity:
			se := &structure.Security{SecurityName: holidays[i].Symbol}
			exist, err := base.Get(base.NewSecurityer(se))
			if err != nil {

			}

			if !exist {

			}

			err = extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{SecurityID: se.ID}, hc))
			if err != nil {

			}

		case structure.HolidaySource:
			src := &structure.Source{Source: holidays[i].Symbol}
			exist, err := base.Get(base.NewSourcer(src))
			if err != nil {

			}

			if !exist {

			}

			err = extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{SourceID: src.ID}, hc))
			if err != nil {

			}

		case structure.HolidaySymbol:
			err = extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{Symbol: holidays[i].Symbol}, hc))
			if err != nil {

			}
		}
	}

	return nil
}