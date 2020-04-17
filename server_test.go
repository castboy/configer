package main

import (
	"configer/server/structure"
	"fmt"
	"testing"
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
	sn, err := GetSourceNameBySymbolName("AUDCAD")
	fmt.Println(sn, err)
}

func Test_DeleteSymbolByName(t *testing.T) {
	err := DeleteSymbolByName("AUDCAD")
	if err != nil {
		fmt.Println(err)
	}
}

func Test_GetConvSymbolInfo(t *testing.T) {
	conv, err := GetConvSymbolInfo(structure.MarginConv, "AUDCAD")
	fmt.Println(conv.ConvSymbol, conv.ConvMultiply, conv.ConvNeed, err)
}

func Test_InsertSymbol(t *testing.T)  {

}

func Test_UpdateSymbol(t *testing.T) {

}

func Test_InsertSource(t *testing.T) {

}

func Test_UpdateSource(t *testing.T) {

}

func Test_GetSources(t *testing.T)  {

}

func Test_GetSourceByName(t *testing.T)  {

}

func Test_GetSourceIDByName(t *testing.T)  {

}

func Test_GetSymbolsBySourceName(t *testing.T)  {

}


func Test_ExportSessions(t *testing.T)  {

}

func Test_SetSession(t *testing.T)  {

}


func Test_GetHolidays(t *testing.T)  {

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

func Test_SetDST(t *testing.T)  {

}

func Test_ExportMarketDST(t *testing.T) {

}

