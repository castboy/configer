package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerA struct {
	a *structure.A
	*xorm.Engine
}

var tablerA *TablerA

func NewTablerA(a *structure.A) *TablerA {
	return &TablerA{
		a,
		xEngine,
	}
}

// implement Tableor
func (t *TablerA) Insert() (int64, error) {
	return t.Table(t.a).Insert(t.a)
}

func (t *TablerA) Delete() (int64, error) {
	return t.Table(t.a).Delete(t.a)
}

func (t *TablerA) Update() (int64, error) {
	return t.Table(t.a).Update(t.a)
}

func (t *TablerA) Get() (bool, error) {
	return t.Table(t.a).Where("name=?", t.a.Name).NoAutoCondition(true).Get(t.a)
}
