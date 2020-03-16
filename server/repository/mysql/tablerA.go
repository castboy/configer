package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerA struct {
	*xorm.Engine
}

var tablerA *TablerA

func GetTablerA() *TablerA {
	if tablerA == nil {
		tablerA = &TablerA{xEngine}
	}

	return tablerA
}

// implement Tableor
func (t *TablerA) Insert(a *structure.A) (int64, error) {
	return t.Table(a).Insert(a)
}

func (t *TablerA) Delete(a *structure.A) (int64, error) {
	return t.Table(a).Delete(a)
}

func (t *TablerA) Update(a *structure.A) (int64, error) {
	return t.Table(a).Update(a)
}

func (t *TablerA) Get(a *structure.A) (bool, error) {
	return t.Table(a).Where("name=?", a.Name).Get(a)
}
