package main

import (
	"configer/server/base"
	"configer/server/constant"
	"configer/server/structure"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
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
	fmt.Println(err, symb.ID)
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
	src, err := GetSourceByName("AUDCAD")
	fmt.Println(src, err)
}

func Test_GetSourceIDByName(t *testing.T) {
	id, err := GetSourceIDByName("AUDCAD")
	fmt.Println(id, err)
}

func Test_GetSymbolsBySourceName(t *testing.T) {
	symbs := GetSymbolsBySourceName("AUDCAD")
	fmt.Println(symbs)
}

func Test_ExportSessions(t *testing.T) {
	ss, err := ExportSessions("AUDCAD", structure.DST, structure.Quote)
	fmt.Println(ss, err)
}

func Test_SetSession(t *testing.T) {
	se := &structure.Session{
		SourceID: 1,
		Type:     structure.Quote,
		Dst:      structure.DST,
		Session: map[int32][]string{
			0: []string{"00:00-01:00"},
		},
	}

	err := SetSession(se)
	fmt.Println(err)
}

func Test_GetHolidays(t *testing.T) {
	hs, err := GetHolidays()
	fmt.Println(hs, err)
}

func Test_InsertHoliday(t *testing.T) {
	ho := &structure.Holiday{
		Enable:      true,
		Date:        "2020-04-10",
		From:        "00:00:00",
		To:          "10:00:00",
		Category:    structure.HolidayAll,
		Description: "just test",
	}

	err := InsertHoliday(ho)
	fmt.Println(err)
}

func Test_UpdateHolidayByID(t *testing.T) {
	ho := &structure.Holiday{
		Enable:      true,
		Date:        "2020-05-10",
		From:        "01:00:01",
		To:          "10:00:00",
		Category:    structure.HolidaySymbol,
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

func Test_SetDST(t *testing.T) {
	err := SetDST(&structure.MarketDST{MarketOwnerType: 0, DST: 2})
	fmt.Println(err)
}

func Test_ExportMarketDST(t *testing.T) {
	md := ExportMarketDST()
	fmt.Println(md)
}

func Test_IsQuotable(t *testing.T) {
	res, err := GetSymbolInfoByName("AUDCAD")
	if err != nil {
		log.Fatal(err)
	}

	ok := IsQuotable(res.Symbol)
	fmt.Println(ok)
}

func Test_IsTradable(t *testing.T) {
	res, err := GetSymbolInfoByName("AUDCAD")
	if err != nil {
		log.Fatal(err)
	}

	ok := IsTradable(res.Symbol)
	fmt.Println(ok)
}

func Test_GetGroupIDByName(t *testing.T) {
	id, err := GetGroupIDByName("TBRD-M00A_EVO")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func Test_FindGroupById(t *testing.T) {
	name, err := FindGroupById(110)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}

func Test_FindGroups(t *testing.T) {
	groups, err := FindGroups()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(groups)
}

func Test_AddGroup(t *testing.T) {
	group := structure.AccountGroup{
		ID:              1000,
		Name:            "testing",
		DepositCurrency: "USD",
		MarginStopOut:   decimal.NewFromFloat(50),
		MarginMode:      constant.LargerLeg,
		MarginCall:      decimal.Zero,
		TradeType:       constant.ABook,
		IsChargeSwap:    false,
		Comment:         "good",
	}
	err := AddGroup(&group)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_ModifyGroup(t *testing.T) {
	group := structure.AccountGroup{
		ID:              1000,
		Name:            "testing",
		DepositCurrency: "USD123123123",
		MarginStopOut:   decimal.NewFromFloat(40),
		MarginMode:      constant.LargerLeg,
		MarginCall:      decimal.Zero,
		TradeType:       constant.ABook,
		IsChargeSwap:    false,
		Comment:         "good",
	}
	err := ModifyGroup(&group)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_DeleteGroupByID(t *testing.T) {

	err := DeleteGroupByID(1000)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_RemoveGroupByName(t *testing.T) {

	err := RemoveGroupByName("testing")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

//conGroupSecurity
func Test_AddConGroupSec(t *testing.T) {
	gs := &structure.ConGroupSec{
		GroupId:        110,
		SecurityId:     1,
		EnableSecurity: true,
		EnableTrade:    true,
		LotMax:         decimal.NewFromFloat(0.01),
		LotMin:         decimal.NewFromFloat(50),
		LotStep:        decimal.NewFromFloat(0.1),
		SpreadDiff:     -2,
		Commission:     decimal.NewFromFloat(1),
	}
	err := AddConGroupSec(gs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_FindConGroupSecById(t *testing.T) {
	cg, err := FindConGroupSecById(3, 24)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
}

func Test_ModifyConGroupSec(t *testing.T) {
	gs := &structure.ConGroupSec{
		ID:             5741,
		GroupId:        110,
		SecurityId:     1,
		EnableSecurity: true,
		EnableTrade:    true,
		LotMax:         decimal.NewFromFloat(0.01),
		LotMin:         decimal.NewFromFloat(50),
		LotStep:        decimal.NewFromFloat(0.1),
		SpreadDiff:     -9,
		Commission:     decimal.NewFromFloat(1),
	}
	err := ModifyConGroupSec(gs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_RemoveConGroupSecById(t *testing.T) {
	err := RemoveConGroupSecById(1100, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_FindAllConGroup(t *testing.T) {
	cgss, err := FindAllConGroup()
	if err != nil {
		t.Fatal(err)
	}
	for i := range cgss {
		fmt.Printf(":%+v\n", cgss[i])
	}
}

func Test_FindConGroupSecsByGroupID(t *testing.T) {
	cgss, err := FindConGroupSecsByGroupID(110)
	if err != nil {
		t.Fatal(err)
	}
	for i := range cgss {
		fmt.Printf(":%+v\n", cgss[i])
	}
}

func Test_IsGroupHoldSecurity(t *testing.T) {
	isHold, err := IsGroupHoldSecurity(1100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isHold)
}
