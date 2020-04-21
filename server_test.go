package main

import (
	"configer/server/base"
	"configer/server/structure"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

func init() {
	Start()
}

func Test_GetSymbolInfoByName(t *testing.T) {
	i, err := GetSymbolInfoByName("AUDCAD")
	fmt.Println(i, err)
}

func Test_GetSymbols(t *testing.T) {
	sbs, err := GetSymbols()
	fmt.Println(sbs, err)
}

func Test_GetSourceNameBySymbolName(t *testing.T) {
	sn, err := GetSourceNameBySymbolName("HK50_200")
	fmt.Println(sn, err)
}

func Test_DeleteSymbolByName(t *testing.T) {
	err := DeleteSymbolByName("wmg")
	if err != nil {
		fmt.Println(err)
	}
}

func Test_GetConvSymbolInfo(t *testing.T) {
	conv, err := GetConvSymbolInfo(structure.MarginConv, "AUDCAD")
	fmt.Println(conv.ConvSymbol, conv.ConvMultiply, conv.ConvNeed, err)
}

func Test_InsertSymbol(t *testing.T)  {
	symb := &structure.Symbol{
		Symbol: "wwwmwg",
		SourceID: 1,
		Leverage: 100,
		SecurityID: 1,
		MarginInitial: decimal.New(1, 1),
		MarginDivider: decimal.New(1, 1),
		Percentage: decimal.New(1, 1),
		Status: structure.SymbolStatus(1),
	}

	err := InsertSymbol(symb)
	fmt.Println(err, symb.ID)
}

func Test_UpdateSymbol(t *testing.T) {
	symb := &structure.Symbol{
		Symbol: "www",
		SourceID: 2,
		Leverage: 100,
		SecurityID: 1,
		MarginInitial: decimal.New(1, 1),
		MarginDivider: decimal.New(1, 1),
		Percentage: decimal.New(1, 1),
		Status: structure.SymbolStatus(1),
	}

	err := UpdateSymbol(symb)
	fmt.Println(err)
}

func Test_InsertSource(t *testing.T) {
	src := &structure.Source{
		Source:          "wmgz",
		SourceCN:        "wmg" ,
		SourceType:      structure.SourceType(1),
		Digits:          1,
		Currency:        "123",
		ContractSize: decimal.NewFromFloat(1),
		StopsLevel: 100,
		ProfitCurrency:  "123",
		MarginCurrency:  "123",
		SwapCurrency:    "123",
		Swap3Day:        time.Weekday(1),
		SwapLong: decimal.NewFromFloat(1),
	}

	err := base.Insert(base.NewSourcer(src))
	fmt.Println(err)
}

func Test_UpdateSource(t *testing.T) {
	src := &structure.Source{
		Source:          "wmgz",
		SourceCN:        "wmgxxx" ,
		SourceType:      structure.SourceType(1),
		Digits:          1,
		Currency:        "123",
		ContractSize: decimal.NewFromFloat(1),
		StopsLevel: 200,
		ProfitCurrency:  "123",
		MarginCurrency:  "123",
		SwapCurrency:    "123",
		Swap3Day:        time.Weekday(1),
		SwapLong: decimal.NewFromFloat(1),
	}

	err := base.Insert(base.NewSourcer(src))
	fmt.Println(err)
}

func Test_GetSources(t *testing.T)  {
	srcs := GetSources()
	fmt.Println(srcs)
}

func Test_GetSourceByName(t *testing.T)  {
	src, err := GetSourceByName("AUDCAD")
	fmt.Println(src, err)
}

func Test_GetSourceIDByName(t *testing.T)  {
	id, err := GetSourceIDByName("AUDCAD")
	fmt.Println(id, err)
}

func Test_GetSymbolsBySourceName(t *testing.T)  {
	symbs := GetSymbolsBySourceName("AUDCAD")
	fmt.Println(symbs)
}


func Test_ExportSessions(t *testing.T)  {
	ss, err := ExportSessions("AUDCAD", structure.DST, structure.Quote)
	fmt.Println(ss, err)
}

func Test_SetSession(t *testing.T)  {
	se := &structure.Session{
		SourceID: 1,
		Type: structure.Quote,
		Dst: structure.DST,
		Session: map[int32][]string{
			0: []string{"00:00-01:00"},
		},
	}

	err := SetSession(se)
	fmt.Println(err)
}


func Test_GetHolidays(t *testing.T)  {
	hs, err := GetHolidays()
	fmt.Println(hs, err)
}

func Test_InsertHoliday(t *testing.T) {
	ho := &structure.Holiday{
		Enable: true,
		Date: "2020-04-10",
		From: "00:00:00",
		To: "10:00:00",
		Category: structure.HolidayAll,
		Description: "just test",
	}

	err := InsertHoliday(ho)
	fmt.Println(err)
}

func Test_UpdateHolidayByID(t *testing.T) {
	ho := &structure.Holiday{
		Enable: true,
		Date: "2020-05-10",
		From: "01:00:01",
		To: "10:00:00",
		Category: structure.HolidaySymbol,
		Description: "just test",
	}

	err := UpdateHolidayByID(2, ho)
	fmt.Println(err)
}

func Test_DeleteHolidayByID(t *testing.T) {
	err := DeleteHolidayByID(2)
	fmt.Println(err)
}


func Test_GetSecurityInfo(t *testing.T) {
	sec, err := GetSecurityInfo(1)
	fmt.Println(sec, err)
}

func Test_GetSecurityNameByID(t *testing.T) {
	name, err := GetSecurityNameByID(1)
	fmt.Println(name, err)
}

func Test_GetAllSecuritiesInfos(t *testing.T) {
	res, err := GetAllSecuritiesInfos()
	if err == nil {
		for i := range res {
			fmt.Println(res[i].SecurityName, res[i].Symbols)
		}
	}
}

func Test_InsertSecurityInfo(t *testing.T) {
	// no usage
}

func Test_DeleteSecurityInfo(t *testing.T) {
	// no usage
}

func Test_GetDST(t *testing.T) {
	dst, err := GetDST(structure.NewYork)
	fmt.Println(dst, err)
}

func Test_SetDST(t *testing.T)  {
	err := SetDST(&structure.MarketDST{MarketOwnerType: 0, DST: 2})
	fmt.Println(err)
}

func Test_ExportMarketDST(t *testing.T) {
	md := ExportMarketDST()
	fmt.Println(md)
}

