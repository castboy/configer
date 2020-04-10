package utils

import "xorm.io/core"

// implement core.IMMaper.

type ObjTable struct {
	Obj,
	Table string
}

func NewObjTable(obj, table string) *ObjTable {
	return &ObjTable{
		Obj:   obj,
		Table: table,
	}
}

type ObjTables []*ObjTable

func SetObjTables(ot ...*ObjTable) (ots ObjTables) {
	ots = ObjTables(ot)
	return
}

func (mm ObjTables) Obj2Table(name string) string {
	for i := range mm {
		if mm[i].Obj == name {
			return mm[i].Table
		}
	}

	return core.SnakeMapper{}.Obj2Table(name)
}

func (mm ObjTables) Table2Obj(name string) string {
	for i := range mm {
		if mm[i].Table == name {
			return mm[i].Obj
		}
	}

	return core.SnakeMapper{}.Table2Obj(name)
}
