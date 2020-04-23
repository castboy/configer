package main

import (
	"configer/server/base"
	"configer/server/extend"
	"configer/server/structure"
	"configer/server/utils"
)

type ExportSymbol struct {
	*structure.Symbol
	*structure.Source
}

func GetSymbolInfoByName(symb string) (*ExportSymbol, error) {
	symbol := &structure.Symbol{Symbol: symb}
	symboler := base.NewSymboler(symbol)
	i, err := base.Get(symboler)
	if err != nil {
		return nil, err
	}

	symbol = i.(*structure.Symbol)

	source := &structure.Source{ID: symbol.SourceID}
	sourcer := base.NewSourcer(source)
	j, err := base.Get(sourcer)
	if err != nil {
		return nil, err
	}

	source = j.(*structure.Source)

	return &ExportSymbol{
		symbol,
		source,
	}, nil
}

func GetSymbols() (symbols []ExportSymbol, err error) {
	symbol := &structure.Symbol{}
	symboler := base.NewSymboler(symbol)
	i, err := base.Export(symboler)
	if err != nil {
		return nil, err
	}

	sbs := i.(map[string]structure.NameIDor)
	for i := range sbs {
		sb := sbs[i].(*structure.Symbol)

		source := &structure.Source{ID: sb.SourceID}
		sourcer := base.NewSourcer(source)
		j, err := base.Get(sourcer)
		if err != nil {
			continue
		}

		source = j.(*structure.Source)

		symbols = append(symbols, ExportSymbol{
			symbol,
			source,
		})
	}

	return
}

func GetFullSymbolName(name string, leverage int32) (fn string, err error) {
	fsn := &structure.FullSymbolName{Sl: structure.SymbolLeverage{Symbol: name, Leverage: leverage}}
	i, err := extend.Get(extend.NewFullSymbolNamer(fsn))
	if err != nil {
		return
	}

	fn = i.(*structure.FullSymbolName).FullName

	return
}

func GetSourceNameBySymbolName(symbolName string) (string, error) {
	symbol := &structure.Symbol{Symbol: symbolName}
	symboler := base.NewSymboler(symbol)
	i, err := base.Get(symboler)
	if err != nil {
		return "", err
	}

	symbol = i.(*structure.Symbol)

	source := &structure.Source{ID: symbol.SourceID}
	sourcer := base.NewSourcer(source)
	j, err := base.Get(sourcer)
	if err != nil {
		return "", err
	}

	source = j.(*structure.Source)

	return source.Source, nil
}

func InsertSymbol(symbol *structure.Symbol) error {
	// insert symbol.
	_, err := base.Get(base.NewSourcer(&structure.Source{ID: symbol.SourceID}))
	if err != nil {
		return err
	}

	_, err = base.Get(base.NewSecurityer(&structure.Security{ID: symbol.SecurityID}))
	if err != nil {
		return err
	}

	symboler := base.NewSymboler(symbol)
	err = base.Insert(symboler)
	if err != nil {
		return err
	}

	// insert full_symbol_name
	fsn := &structure.FullSymbolName{}
	fsn.Sl.Symbol = utils.GetRequestSymbol(symbol.Symbol)
	fsn.Sl.Leverage = symbol.Leverage
	fsn.FullName = symbol.Symbol

	return extend.Insert(extend.NewFullSymbolNamer(fsn))
}

func UpdateSymbol(symbol *structure.Symbol) error {
	symboler := base.NewSymboler(symbol)
	_, err := base.Update(symboler)

	// effect for full_symbol_name ?
	// TODO

	return err
}

func DeleteSymbolByName(symbolName string) error {
	symboler := base.NewSymboler(&structure.Symbol{Symbol: symbolName})
	_, err := base.Delete(symboler)
	return err
}

func GetConvSymbolInfo(t structure.ConvType, symbolName string) (*structure.ConvInfo, error) {
	symbol := &structure.Symbol{Symbol: symbolName}
	symboler := base.NewSymboler(symbol)
	i, err := base.Get(symboler)
	if err != nil {
		return nil, err
	}

	symbol = i.(*structure.Symbol)

	convSymb := &structure.ConvSymbol{ConvType: t, SourceID: symbol.SourceID}
	j, err := extend.Get(extend.NewConvSymboler(convSymb))
	if err != nil {
		return nil, err
	}

	return j.(*structure.ConvSymbol).ConvInfo, nil
}

func InsertSource(source *structure.Source) error {
	// insert source
	err := base.Insert(base.NewSourcer(source))
	if err != nil {
		return err
	}

	// insert conv_symbol_info
	sources := GetSources()

	srcs := []structure.Source{}
	for j := range sources {
		srcs = append(srcs, *sources[j])
	}

	mc := utils.BuildConvInfo(source.MarginCurrency, srcs)
	cs := &structure.ConvSymbol{ConvType: structure.MarginConv, SourceID: source.ID, ConvInfo: mc}
	err = extend.Insert(extend.NewConvSymboler(cs))
	if err != nil {
		return err
	}

	pc := utils.BuildConvInfo(source.ProfitCurrency, srcs)
	cs = &structure.ConvSymbol{ConvType: structure.ProfitConv, SourceID: source.ID, ConvInfo: pc}
	err = extend.Insert(extend.NewConvSymboler(cs))
	if err != nil {
		return err
	}

	return nil
}

func UpdateSource(source *structure.Source) error {
	src := &structure.Source{Source: source.Source}
	i, err := base.Get(base.NewSourcer(src))
	if err != nil {
		return err
	}

	src = i.(*structure.Source)

	// no updated fields.
	source.ID = src.ID
	source.Source = src.Source
	source.SourceCN = src.SourceCN
	source.ProfitMode = src.ProfitMode
	source.SwapCurrency = src.SwapCurrency

	_, err = base.Update(base.NewSourcer(source))
	return err
}

func GetSources() (sources map[int]*structure.Source) {
	sources = map[int]*structure.Source{}

	i, err := base.Export(base.NewSourcer(&structure.Source{}))
	if err != nil {
		return nil
	}

	srcs := i.(map[string]structure.NameIDor)
	for i := range srcs {
		src := srcs[i].(*structure.Source)
		sources[src.ID] = src
	}

	return
}

func GetSourceByName(sourceName string) (*structure.Source, error) {
	i, err := base.Get(base.NewSourcer(&structure.Source{Source: sourceName}))
	if err != nil {
		return nil, err
	}

	return i.(*structure.Source), nil

}

func GetSourceIDByName(sourceName string) (ID int, err error) {
	src, err := GetSourceByName(sourceName)
	if err != nil {
		return
	}

	return src.ID, nil
}

func GetSymbolsBySourceName(sourceName string) (symbols []string) {
	id, err := GetSourceIDByName(sourceName)
	if err != nil {
		return nil
	}

	i, err := base.Export(base.NewSymboler(&structure.Symbol{}))
	if err != nil {
		return
	}

	symbs := i.(map[string]structure.NameIDor)
	for i := range symbs {
		symb := symbs[i].(*structure.Symbol)
		if symb.SourceID == id {
			symbols = append(symbols, symb.Symbol)
		}
	}

	return
}

func GetDST(marketType structure.MarketType) (dst structure.DSTType, err error) {
	i, err := base.Get(base.NewMarketDSTer(&structure.MarketDST{MarketOwnerType: marketType}))
	if err != nil {
		return
	}

	return i.(*structure.MarketDST).DST, nil
}

func SetDST(marketDST *structure.MarketDST) error {
	_, err := base.Update(base.NewMarketDSTer(marketDST))
	return err
}

func ExportMarketDST() map[structure.MarketType]structure.DSTType {
	i, _ := base.Export(base.NewMarketDSTer(&structure.MarketDST{}))
	md := i.(map[structure.MarketType]*structure.MarketDST)

	res := map[structure.MarketType]structure.DSTType{}
	for j := range md {
		res[md[j].MarketOwnerType] = md[j].DST
	}

	return res
}
