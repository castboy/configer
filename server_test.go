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

func Test_InsertSymbol(t *testing.T) {
	symb := &structure.Symbol{
		Symbol:        "www",
		SourceID:      1,
		Leverage:      100,
		SecurityID:    1,
		MarginInitial: decimal.New(1, 1),
		MarginDivider: decimal.New(1, 1),
		Percentage:    decimal.New(1, 1),
		Status:        structure.SymbolStatus(1),
	}

	err := InsertSymbol(symb)
	fmt.Println(err)
}

func Test_UpdateSymbol(t *testing.T) {
	symb := &structure.Symbol{
		Symbol:        "www",
		SourceID:      2,
		Leverage:      100,
		SecurityID:    1,
		MarginInitial: decimal.New(1, 1),
		MarginDivider: decimal.New(1, 1),
		Percentage:    decimal.New(1, 1),
		Status:        structure.SymbolStatus(1),
	}

	err := UpdateSymbol(symb)
	fmt.Println(err)
}

func Test_InsertSource(t *testing.T) {
	src := &structure.Source{
		Source:         "wmgz",
		SourceCN:       "wmg",
		SourceType:     structure.SourceType(1),
		Digits:         1,
		Currency:       "123",
		ContractSize:   decimal.NewFromFloat(1),
		StopsLevel:     100,
		ProfitCurrency: "123",
		MarginCurrency: "123",
		SwapCurrency:   "123",
		Swap3Day:       time.Weekday(1),
		SwapLong:       decimal.NewFromFloat(1),
	}

	err := base.Insert(base.NewSourcer(src))
	fmt.Println(err)
}

func Test_UpdateSource(t *testing.T) {
	src := &structure.Source{
		Source:         "wmgz",
		SourceCN:       "wmgxxx",
		SourceType:     structure.SourceType(1),
		Digits:         1,
		Currency:       "123",
		ContractSize:   decimal.NewFromFloat(1),
		StopsLevel:     200,
		ProfitCurrency: "123",
		MarginCurrency: "123",
		SwapCurrency:   "123",
		Swap3Day:       time.Weekday(1),
		SwapLong:       decimal.NewFromFloat(1),
	}

	err := base.Insert(base.NewSourcer(src))
	fmt.Println(err)
}

func Test_GetSources(t *testing.T) {
	srcs := GetSources()
	fmt.Println(srcs)
}

func Test_GetSourceByName(t *testing.T) {

}

func Test_GetSourceIDByName(t *testing.T) {

}

func Test_GetSymbolsBySourceName(t *testing.T) {

}

func Test_ExportSessions(t *testing.T) {

}

func Test_SetSession(t *testing.T) {

}

func Test_GetHolidays(t *testing.T) {

}

func Test_GetHolidayByID(t *testing.T) {

}

func Test_InsertHoliday(t *testing.T) {

}

func Test_UpdateHolidayByID(t *testing.T) {

}

func Test_DeleteHolidayByID(t *testing.T) {

}

func Test_GetSecurityInfo(t *testing.T) {

}

func Test_GetSecurityNameByID(t *testing.T) {

}

func Test_GetAllSecuritiesInfos(t *testing.T) {

}

func Test_UpdateSecurityInfo(t *testing.T) {

}

func Test_InsertSecurityInfo(t *testing.T) {

}

func Test_DeleteSecurityInfo(t *testing.T) {

}

func Test_GetDST(t *testing.T) {

}

func Test_SetDST(t *testing.T) {

}

func Test_ExportMarketDST(t *testing.T) {

}
