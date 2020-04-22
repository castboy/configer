package main

import (
	"configer/server/base"
	"configer/server/extend"
	"configer/server/structure"
)

func Start() (err error) {
	// cache group.
	group := &structure.AccountGroup{}
	grouper := base.NewGrouper(group)
	err = base.Cache(grouper)
	if err != nil {

	}
	// cache symbol.
	symbol := &structure.Symbol{}
	symboler := base.NewSymboler(symbol)
	err = base.Cache(symboler)
	if err != nil {

	}

	// cache source.
	source := &structure.Source{}
	sourcer := base.NewSourcer(source)
	err = base.Cache(sourcer)
	if err != nil {

	}

	// cache session.
	for i := 0; i < int(structure.SessionTypeLength); i++ {
		for j := 0; j < int(structure.DSTTypeLength); j++ {
			sess := &structure.Session{Type: structure.SessionType(i), Dst: structure.DSTType(j)}
			err = base.Cache(base.NewSessioner(sess))
			if err != nil {

			}
		}
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
	err = extend.Cache(extend.NewFullSymbolNamer(fsn))
	if err != nil {

	}

	// cache conv_symbol
	for i := 0; i < int(structure.ConvTypeLength); i++ {
		cs := &structure.ConvSymbol{ConvType: structure.ConvType(i)}
		err = extend.Cache(extend.NewConvSymboler(cs))
		if err != nil {

		}
	}

	// cache holiday
	ho := &structure.Holiday{}
	holidayer := base.NewHolidayer(ho)
	err = base.Cache(holidayer)
	if err != nil {

	}

	i, err := base.Export(holidayer)
	if err != nil {

	}

	hos := i.(map[string]map[int]*structure.Holiday)
	for date := range hos {
		for id := range hos[date] {
			appendSymbolID(hos[date][id])
		}
	}

	return nil
}