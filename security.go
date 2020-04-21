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


