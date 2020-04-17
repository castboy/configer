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
	//
	//i, err := GetSymbolInfoByName("AUDCAD")
	//fmt.Println(i, err)
	//
	//sbs, err := GetSymbols()
	//fmt.Println(sbs, err)
	//
	//sn, err := GetSourceNameBySymbolName("AUDCAD")
	//fmt.Println(sn, err)
	//
	//symb := i.Symbol
	//symb.ID = 0
	//symb.Symbol = "wmxx"
	//
	//err = DeleteSymbolByName(symb.Symbol)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//i, err = GetSymbolInfoByName("wmxx")
	//fmt.Println(i, err)

	//conv, err := GetConvSymbolInfo(structure.MarginConv, "AUDCAD")
	//fmt.Println(conv.ConvSymbol, conv.ConvMultiply, conv.ConvNeed, err)

	//source := &structure.Source{Source: "AUDCAD"}
	//sourcer := base.NewSourcer(source)
	//i, exist, err := base.Get(sourcer)
	//fmt.Println(exist, err)
	//
	//source = i.(*structure.Source)
	//source.Source = "wmgg"
	//source.ID = 0
	//
	//fmt.Println(source)
	//sourcer = base.NewSourcer(source)
	//
	//_, err = base.Insert(sourcer)
	//fmt.Println(err)
	//
	//source.Digits = 6
	//
	//_, err = base.Update(sourcer)
	//fmt.Println(err)
	//
	//i, exist, err = base.Get(sourcer)
	//fmt.Println(i, exist, err)

	//fmt.Println(GetSources())

	//fmt.Println(GetSourceByName("AUDCAD"))

	//fmt.Println(GetSymbolsBySourceName("AUDCAD"))

	//ss, err := ExportSessions("AUDCAD", 1, 0)
	//fmt.Println(ss, err)

	//hs, err := GetHolidays()
	//fmt.Println(hs, err)

	conv, err := GetConvSymbolInfo(structure.MarginConv, "AUDCAD")
	fmt.Println(conv, err)
}

type ExportSymbol struct {
	*structure.Symbol
	*structure.Source
}

func GetSymbolInfoByName(symb string) (*ExportSymbol, error) {
	symbol := &structure.Symbol{Symbol: symb}
	symboler := base.NewSymboler(symbol)
	i, exist, err := base.Get(symboler)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, err
	}

	symbol = i.(*structure.Symbol)

	source := &structure.Source{ID: symbol.SourceID}
	sourcer := base.NewSourcer(source)
	j, exist, err := base.Get(sourcer)
	if err != nil {
		return nil, err
	}

	if !exist {
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
		j, exist, err := base.Get(sourcer)
		if err != nil || !exist {
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

func GetSourceNameBySymbolName(symbolName string) (string, error) {
	symbol := &structure.Symbol{Symbol: symbolName}
	symboler := base.NewSymboler(symbol)
	i, exist, err := base.Get(symboler)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", nil
	}

	symbol = i.(*structure.Symbol)

	source := &structure.Source{ID: symbol.SourceID}
	sourcer := base.NewSourcer(source)
	j, exist, err := base.Get(sourcer)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", nil
	}

	source = j.(*structure.Source)

	return source.Source, nil
}

func InsertSymbol(symbol *structure.Symbol) error {
	symboler := base.NewSymboler(symbol)
	_, err := base.Insert(symboler)
	return err
}

func UpdateSymbol(symbol *structure.Symbol) error {
	symboler := base.NewSymboler(symbol)
	_, err := base.Update(symboler)
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
	i, exist, err := base.Get(symboler)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	symbol = i.(*structure.Symbol)

	convSymb := &structure.ConvSymbol{ConvType: t, SourceID: symbol.SourceID}
	j, exist, err := extend.Get(extend.NewConvSymboler(convSymb))
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	return j.(*structure.ConvSymbol).ConvInfo, nil
}

func InsertSource(source *structure.Source) error {
	_, err := base.Insert(base.NewSourcer(source))
	return err
}

func UpdateSource(source *structure.Source) error {
	_, err := base.Update(base.NewSourcer(source))
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
	i, exist, err := base.Get(base.NewSourcer(&structure.Source{Source: sourceName}))
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, nil
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
	fmt.Println(id)
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
	i, exist, err := base.Get(base.NewHolidayer(&structure.Holiday{ID: ID}))
	if err != nil {
		return
	}

	if !exist {
		return
	}

	return i.(*structure.Holiday), nil
}

func InsertHoliday(holi *structure.Holiday) error {
	_, err := base.Insert(base.NewHolidayer(holi))
	return err
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

func GetSecurityInfo(id int) (bean *structure.Security, err error) {
	i, exist, err := base.Get(base.NewSecurityer(&structure.Security{ID: id}))
	if err != nil {
		return
	}

	if !exist {
		return nil, nil
	}

	return i.(*structure.Security), nil
}

func GetSecurityNameByID(id int) (s string, err error) {
	sec, err := GetSecurityInfo(id)
	if err != nil {
		return
	}

	return sec.SecurityName, nil
}

func GetAllSecuritiesInfos() (res []*structure.Security, err error) {
	i, err := base.Export(base.NewSecurityer(&structure.Security{}))
	if err != nil {
		return
	}

	se := i.(map[string]structure.NameIDor)

	for j := range se {
		res = append(res, se[j].(*structure.Security))
	}

	return
}

func UpdateSecurityInfo(id int, info *structure.Security) error {
	info.ID = id
	_, err := base.Update(base.NewSecurityer(info))
	return err
}

func InsertSecurityInfo(info *structure.Security) error {
	_, err := base.Insert(base.NewSecurityer(info))
	return err
}

func DeleteSecurityInfo(id int) error {
	_, err := base.Delete(base.NewSecurityer(&structure.Security{ID: id}))
	return err
}

func GetDST(marketType structure.MarketType) (dst structure.DSTType, err error) {
	i, exist, err := base.Get(base.NewMarketDSTer(&structure.MarketDST{MarketOwnerType: marketType}))
	if err != nil {
		return
	}

	if !exist {
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

func Start() (err error) {
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


func appendSymbolID(ho *structure.Holiday) (err error) {
	switch ho.Category {
	case structure.HolidayAll:

	case structure.HolidaySecurity:
		i, exist, err := base.Get(base.NewSecurityer(&structure.Security{SecurityName: ho.Symbol}))
		if err != nil {
			return
		}

		if !exist {
			return errors.NotFoundf("%v", &structure.Security{SecurityName: ho.Symbol})
		}

		se := i.(*structure.Security)
		ho.SymbolID = se.ID

	case structure.HolidaySource:
		i, exist, err := base.Get(base.NewSourcer(&structure.Source{Source: ho.Symbol}))
		if err != nil {
			return
		}

		if !exist {
			return errors.NotFoundf("%v", &structure.Source{Source: ho.Symbol})
		}

		src := i.(*structure.Source)
		ho.SymbolID = src.ID

	case structure.HolidaySymbol:
		i, exist, err := base.Get(base.NewSymboler(&structure.Symbol{Symbol: ho.Symbol}))
		if err != nil {
			return
		}

		if !exist {
			return errors.NotFoundf("%v", &structure.Symbol{Symbol: ho.Symbol})
		}

		symb := i.(*structure.Symbol)
		ho.SymbolID = symb.ID

	default:
		panic(errors.NotValidf("holiday category: %d", ho.Category))
	}

	return err
}

