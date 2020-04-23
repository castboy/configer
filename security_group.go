package main

import (
	"configer/server/base"
	"configer/server/constant"
	"configer/server/structure"
	"fmt"
)

func GetSecurityInfo(id int) (bean *structure.Security, err error) {
	i, err := base.Get(base.NewSecurityer(&structure.Security{ID: id}))
	if err != nil {
		return
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

	// append symbols
	symbols := map[int][]string{}

	j, err := base.Export(base.NewSymboler(&structure.Symbol{}))
	if err != nil {
		return
	}

	sbs := j.(map[string]structure.NameIDor)
	for j := range sbs {
		sb := sbs[j].(*structure.Symbol)
		symbols[sb.SecurityID] = append(symbols[sb.SecurityID], sb.Symbol)
	}

	for i := range res {
		res[i].Symbols = symbols[res[i].ID]
	}

	return
}

func UpdateSecurityInfo(id int, info *structure.Security) error {
	info.ID = id
	_, err := base.Update(base.NewSecurityer(info))
	return err
}

func InsertSecurityInfo(info *structure.Security) error {
	return base.Insert(base.NewSecurityer(info))
}

func DeleteSecurityInfo(id int) error {
	symbName := []string{}

	i, err := base.Export(base.NewSymboler(&structure.Symbol{}))
	if err != nil {
		return err
	}

	sbs := i.(map[string]structure.NameIDor)
	for j := range sbs {
		sb := sbs[j].(*structure.Symbol)
		if sb.SecurityID == id {
			symbName = append(symbName, sb.Symbol)
		}
	}

	if len(symbName) != 0 {
		return constant.NewErr(constant.ArgsErr, fmt.Sprintf("Forbidden DeleteSecurityInfo, securityID: %d", id))
	}

	_, err = base.Delete(base.NewSecurityer(&structure.Security{ID: id}))

	return err
}

//Group

func GetGroupIDByName(groupName string) (groupID int, err error) {
	group := &structure.AccountGroup{Name: groupName}
	grouper := base.NewGrouper(group)
	i, err := base.Get(grouper)
	if err != nil {
		return -1, err
	}
	group = i.(*structure.AccountGroup)
	group.GetID()
	return group.ID, nil
}

func FindGroupById(groupID int) (*structure.AccountGroup, error) {
	group := &structure.AccountGroup{ID: groupID}
	grouper := base.NewGrouper(group)
	i, err := base.Get(grouper)
	if err != nil {
		return nil, err
	}
	group = i.(*structure.AccountGroup)
	return group, nil
}

func FindGroups() ([]structure.AccountGroup, error) {
	groups := make([]structure.AccountGroup, 0)
	group := &structure.AccountGroup{}
	grouper := base.NewGrouper(group)
	i, err := base.Export(grouper)
	if err != nil {
		return nil, err
	}
	gs := i.(map[string]structure.NameIDor)

	for j := range gs {
		groups = append(groups, *gs[j].(*structure.AccountGroup))
	}

	return groups, nil
}

func AddGroup(ug *structure.AccountGroup) error {
	grouper := base.NewGrouper(ug)
	err := base.Insert(grouper)
	if err != nil {
		return nil
	}
	return nil
}

func ModifyGroup(ug *structure.AccountGroup) error {
	grouper := base.NewGrouper(ug)
	_, err := base.Update(grouper)
	if err != nil {
		return nil
	}
	return nil
}

func DeleteGroupByID(id int) error {
	group := &structure.AccountGroup{ID: id}
	grouper := base.NewGrouper(group)
	_, err := base.Delete(grouper)
	if err != nil {
		return nil
	}
	return nil
}

//TODO: sync cache
func RemoveGroupByName(name string) error {
	group := &structure.AccountGroup{Name: name}
	grouper := base.NewGrouper(group)
	_, err := base.Delete(grouper)
	if err != nil {
		return nil
	}
	return nil
}

//conGroupSecurity
func AddConGroupSec(gs *structure.ConGroupSec) error {
	_, err := FindGroupById(gs.GroupId)
	if err != nil {
		return err
	}

	_, err = GetSecurityInfo(gs.SecurityId)
	if err != nil {
		return err
	}

	err = base.Insert(base.NewConGroupSec(gs))
	if err != nil {
		return err
	}

	return err
}
func FindConGroupSecById(groupId, securityId int) (cg *structure.ConGroupSec, err error) {
	cg = &structure.ConGroupSec{GroupId: groupId, SecurityId: securityId}
	cgser := base.NewConGroupSec(cg)
	i, err := base.Get(cgser)
	if err != nil {
		return nil, err
	}
	cg = i.(*structure.ConGroupSec)

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
	cgs := &structure.ConGroupSec{GroupId: groupId, SecurityId: securityId}
	cgser := base.NewConGroupSec(cgs)
	_, err := base.Delete(cgser)
	return err
}

func FindAllConGroup() (cgss []*structure.ConGroupSec, err error) {
	cgs := &structure.ConGroupSec{}
	cgser := base.NewConGroupSec(cgs)
	i, err := base.Export(cgser)
	if err != nil {
		return nil, err
	}
	cgmap := i.(map[int]map[int]*structure.ConGroupSec)
	for j := range cgmap {
		for k := range cgmap[j] {
			cgss = append(cgss, cgmap[j][k])
		}
	}

	return
}

func FindConGroupSecsByGroupID(groupId int) (cgss []*structure.ConGroupSec, err error) {
	cgs := &structure.ConGroupSec{}
	cgser := base.NewConGroupSec(cgs)
	i, err := base.Export(cgser)
	if err != nil {
		return nil, err
	}
	cgmap := i.(map[int]map[int]*structure.ConGroupSec)
	for j := range cgmap {
		if j != groupId {
			continue
		}
		for k := range cgmap[j] {
			cgss = append(cgss, cgmap[j][k])
		}
	}

	return
}

func IsGroupHoldSecurity(groupID int) (bool, error) {
	cgs := &structure.ConGroupSec{}
	cgser := base.NewConGroupSec(cgs)
	i, err := base.Export(cgser)
	if err != nil {
		return false, err
	}
	cgmap := i.(map[int]map[int]*structure.ConGroupSec)
	if len(cgmap[groupID]) != 0 {
		return true, nil
	}
	return false, nil
}
