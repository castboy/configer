package main

import (
	"configer/server/base"
	"configer/server/extend"
	"configer/server/structure"
	"fmt"
	"github.com/juju/errors"
)

func main() {
	Start()
}

func Start() error {
	// cache symbol.
	//symbol := &indexNameID.Symbol{Symbol: "AUDCAD"}
	//symboler := base.NewSymboler(symbol)
	//err := base.Cache(symboler)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//i, exist, err := base.Get(symboler)
	//fmt.Println(i, exist, err)

	// cache session.
	/*sess := &structIndexID.Session{}
	err = base.Cache(base.NewSessioner(sess))
	if err != nil {

	}*/

	// cache security.
	sec := &structure.Security{}
	err := base.Cache(base.NewSecurityer(sec))
	if err != nil {

	}

	// cache market_dst.
	/*md := &structIndexID.MarketDST{}
	err = base.Cache(base.NewMarketDSTer(md))
	if err != nil {

	}

	// cache full_symbol_name.
	fsn := &structIndexID.FullSymbolName{}
	err = extend.Cache(extend.NewFullSymbolNamer(symbol, fsn))
	if err != nil {

	}

	// cache conv_symbol
	cs := &structIndexID.ConvSymbol{}
	err = extend.Cache(extend.NewConvSymboler(symbol, cs))
	if err != nil {

	}*/

	// cache holiday
	ho := &structure.Holiday{}
	holidayer := base.NewHolidayer(ho)
	err = base.Cache(holidayer)
	if err != nil {
		fmt.Println(err)
	}

	// cache holiday.calc
	hs, err := base.Export(holidayer)
	if err != nil {
		fmt.Println(hs, err)
	}

	holidays := hs.(map[int]structure.IDor)
	for i := range holidays {
		err = CacheHolidayCalc(holidays[i].(*structure.Holiday))
	}

	i, err := extend.Export(extend.NewHolidayCalcer(nil, &structure.HolidayCalc{}))
	fmt.Println(i, err)

	return nil
}


func CacheHolidayCalc(ho *structure.Holiday) (err error) {
	hc := &structure.HolidayCalc{}
	hc.ID = ho.ID
	hc.Date = ho.Date
	hc.TimeSpans = append(hc.TimeSpans, &structure.TimeSpan{From: ho.From, To: ho.To})

	switch ho.Category {
	case structure.HolidayAll:
		err := extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{}, hc))
		if err != nil {

		}

	case structure.HolidaySecurity:
		i, exist, err := base.Get(base.NewSecurityer(&structure.Security{SecurityName: ho.Symbol}))
		if err != nil {

		}

		if !exist {
			fmt.Println(errors.NotFoundf("%v", &structure.Security{SecurityName: ho.Symbol}))
		}

		se := i.(*structure.Security)

		err = extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{SecurityID: se.ID}, hc))
		if err != nil {

		}

	case structure.HolidaySource:
		i, exist, err := base.Get(base.NewSourcer(&structure.Source{Source: ho.Symbol}))
		if err != nil {

		}

		if !exist {

		}

		src := i.(*structure.Source)

		err = extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{SourceID: src.ID}, hc))
		if err != nil {

		}

	case structure.HolidaySymbol:
		err = extend.Cache(extend.NewHolidayCalcer(&structure.Symbol{Symbol: ho.Symbol}, hc))
		if err != nil {

		}
	}
	return err
}

