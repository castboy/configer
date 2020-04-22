package main

import (
	"configer/server/constant"
	"configer/server/structure"
	"fmt"
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
	name,err:=FindGroupById(110)
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
		MarginStopOut: decimal.NewFromFloat(50),
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
		MarginStopOut: decimal.NewFromFloat(40),
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

//conGroupSecurity
func Test_AddConGroupSec(t *testing.T){
	gs:=&structure.ConGroupSec{
		GroupId:110,
		SecurityId:1,
		EnableSecurity:true,
		EnableTrade:true,
		LotMax:decimal.NewFromFloat(0.01),
		LotMin:decimal.NewFromFloat(50),
		LotStep:decimal.NewFromFloat(0.1),
		SpreadDiff:-2,
		Commission:decimal.NewFromFloat(1),
	}
	err:=AddConGroupSec(gs)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_FindConGroupSecById(t *testing.T){
	cg,err:=FindConGroupSecById(3,24)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(cg)
}

func Test_ModifyConGroupSec(t *testing.T){
	gs:=&structure.ConGroupSec{
		ID:5741,
		GroupId:110,
		SecurityId:1,
		EnableSecurity:true,
		EnableTrade:true,
		LotMax:decimal.NewFromFloat(0.01),
		LotMin:decimal.NewFromFloat(50),
		LotStep:decimal.NewFromFloat(0.1),
		SpreadDiff:-9,
		Commission:decimal.NewFromFloat(1),
	}
	err:=ModifyConGroupSec(gs)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_RemoveConGroupSecById(t *testing.T){
	err:=RemoveConGroupSecById(1100,1)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("success")
}

func Test_FindAllConGroup(t *testing.T){
	cgss,err:=FindAllConGroup()
	if err!=nil{
		t.Fatal(err)
	}
	for i:=range cgss {
		fmt.Printf(":%+v\n",cgss[i])
	}
}

func Test_FindConGroupSecsByGroupID(t *testing.T){
	cgss,err:=FindConGroupSecsByGroupID(110)
	if err!=nil{
		t.Fatal(err)
	}
	for i:=range cgss {
		fmt.Printf(":%+v\n",cgss[i])
	}
}

func Test_IsGroupHoldSecurity(t *testing.T){
	isHold,err:=IsGroupHoldSecurity(1100)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(isHold)
}