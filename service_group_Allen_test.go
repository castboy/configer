package main

import (
	"configer/server/constant"
	"configer/server/structure"
	"github.com/shopspring/decimal"
	"testing"
)

func Test_GetGroupIDByName(t *testing.T){
	id,err:=GetGroupIDByName("TBRD-M00A_EVO")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
}

func Test_FindGroupById(t *testing.T){
	name,err:=FindGroupById(3)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(name)
}

func Test_FindGroups(t *testing.T){
	groups,err:=FindGroups()
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(groups)
}

func Test_AddGroup(t *testing.T){
	group:=structure.AccountGroup{
		ID:1000,
		Name:"testing",
		DepositCurrency:"USD",
		MarginStopOut: decimal.Zero,
		MarginMode:constant.LargerLeg,
		MarginCall:decimal.Zero,
		TradeType:constant.ABook,
		IsChargeSwap:false,
		Comment:"good",
	}
	err:=AddGroup(&group)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_ModifyGroup(t *testing.T){
	group:=structure.AccountGroup{
		ID:1000,
		Name:"testing",
		DepositCurrency:"USD123123123",
		MarginStopOut: decimal.Zero,
		MarginMode:constant.LargerLeg,
		MarginCall:decimal.Zero,
		TradeType:constant.ABook,
		IsChargeSwap:false,
		Comment:"good",
	}
	err:=ModifyGroup(&group)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_DeleteGroupByID(t *testing.T){

	err:=DeleteGroupByID(1000)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_RemoveGroupByName(t *testing.T){

	err:=RemoveGroupByName("testing")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}