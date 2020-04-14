package main

import (
	"configer/server/base"
	"configer/server/extend"
	"configer/server/structure"
	structIndexID "configer/server/structure/indexID"
	structIndexNameID "configer/server/structure/indexNameID"
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
	sec := &structIndexNameID.Security{}
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
	ho := &structIndexID.Holiday{}
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

	holidays := hs.(map[int]structIndexID.IDor)
	for i := range holidays {
		err = CacheHolidayCalc(holidays[i].(*structIndexID.Holiday))
	}

	i, err := extend.Export(extend.NewHolidayCalcer(nil, &structure.HolidayCalc{}))
	fmt.Println(i, err)

	return nil
}


func CacheHolidayCalc(ho *structIndexID.Holiday) (err error) {
	hc := &structure.HolidayCalc{}
	hc.ID = ho.ID
	hc.Date = ho.Date
	hc.TimeSpans = append(hc.TimeSpans, &structure.TimeSpan{From: ho.From, To: ho.To})

	switch ho.Category {
	case structIndexID.HolidayAll:
		err := extend.Cache(extend.NewHolidayCalcer(&structIndexNameID.Symbol{}, hc))
		if err != nil {

		}

	case structIndexID.HolidaySecurity:
		i, exist, err := base.Get(base.NewSecurityer(&structIndexNameID.Security{SecurityName: ho.Symbol}))
		if err != nil {

		}

		if !exist {
			fmt.Println(errors.NotFoundf("%v", &structIndexNameID.Security{SecurityName: ho.Symbol}))
		}

		se := i.(*structIndexNameID.Security)

		err = extend.Cache(extend.NewHolidayCalcer(&structIndexNameID.Symbol{SecurityID: se.ID}, hc))
		if err != nil {

		}

	case structIndexID.HolidaySource:
		i, exist, err := base.Get(base.NewSourcer(&structIndexNameID.Source{Source: ho.Symbol}))
		if err != nil {

		}

		if !exist {

		}

		src := i.(*structIndexNameID.Source)

		err = extend.Cache(extend.NewHolidayCalcer(&structIndexNameID.Symbol{SourceID: src.ID}, hc))
		if err != nil {

		}

	case structIndexID.HolidaySymbol:
		err = extend.Cache(extend.NewHolidayCalcer(&structIndexNameID.Symbol{Symbol: ho.Symbol}, hc))
		if err != nil {

		}
	}
	return err
}

