package main

import (
	"configer/server/base"
	"configer/server/structure"
)

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
