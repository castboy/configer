package main

import (
	"configer/server/base"
	"configer/server/structure"
	"fmt"
	"github.com/juju/errors"
)

func AddConGroupSec(gs *structure.ConGroupSec) error {
	group,err:=FindGroupById(gs.GroupId)
	if err!=nil || group==nil{
		fmt.Printf(":%+v",err)
		return errors.NotValidf("group id: %d", gs.GroupId)
	}
	sec,err := GetSecurityInfo(gs.SecurityId)
	if sec==nil || err != nil {
		return err
	}

	err = base.Insert(base.NewConGroupSec(gs))
	if err != nil {
		return err
	}

	return err
}
func FindConGroupSecById(groupId, securityId int) (cg *structure.ConGroupSec, err error) {
	cg= &structure.ConGroupSec{GroupId:groupId,SecurityId:securityId}
	cgser := base.NewConGroupSec(cg)
	i, err := base.Get(cgser)
	if err!=nil{
		return nil,err
	}
	cg=i.(*structure.ConGroupSec)

	return
}

func ModifyConGroupSec(cg *structure.ConGroupSec) error {
	cgser := base.NewConGroupSec(cg)
	_, err := base.Update(cgser)
	if err != nil {
		return err
	}
	return err
}

//TODO: sync cache

func RemoveConGroupSecById(groupId, securityId int) error {
	cgs:= &structure.ConGroupSec{GroupId:groupId,SecurityId:securityId}
	cgser := base.NewConGroupSec(cgs)
	_,err:=base.Delete(cgser)
	return err
}

func FindAllConGroup() (cgss []*structure.ConGroupSec, err error) {
	cgs:= &structure.ConGroupSec{}
	cgser := base.NewConGroupSec(cgs)
	i,err:=base.Export(cgser)
	if err!=nil{
		return nil,err
	}
	cgmap:=i.(map[int]map[int]*structure.ConGroupSec)
	for j:=range cgmap{
		for k:=range cgmap[j]{
			cgss=append(cgss,cgmap[j][k])
		}
	}

	return
}

func FindConGroupSecsByGroupID(groupId int) (cgss []*structure.ConGroupSec, err error) {
	cgs:= &structure.ConGroupSec{}
	cgser := base.NewConGroupSec(cgs)
	i,err:=base.Export(cgser)
	if err!=nil{
		return nil,err
	}
	cgmap:=i.(map[int]map[int]*structure.ConGroupSec)
	for j:=range cgmap{
		if j!=groupId {
			continue
		}
		for k:=range cgmap[j]{
			cgss=append(cgss,cgmap[j][k])
		}
	}

	return
}

func IsGroupHoldSecurity(groupID int) (bool, error) {
	cgs:= &structure.ConGroupSec{}
	cgser := base.NewConGroupSec(cgs)
	i,err:=base.Export(cgser)
	if err!=nil{
		return false,err
	}
	cgmap:=i.(map[int]map[int]*structure.ConGroupSec)
	if cgmap[groupID]!=nil || len(cgmap[groupID])!=0{
		return true,nil
	}
	return false,nil
}