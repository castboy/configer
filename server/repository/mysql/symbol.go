package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerSymbol struct {
	a *structure.Symbol
	*xorm.Engine
}

var tablerSymbol *TablerSymbol

func NewTablerSymbol(a *structure.Symbol) *TablerSymbol {
	return &TablerSymbol{
		a,
		xEngine,
	}
}

// implement Tableor
func (t *TablerSymbol) Insert() (int64, error) {
	return t.Table(t.a).Insert(t.a)
}

func (t *TablerSymbol) Delete() (int64, error) {
	return t.Table(t.a).Delete(t.a)
}

func (t *TablerSymbol) Update() (int64, error) {
	return t.Table(t.a).Update(t.a)
}

func (t *TablerSymbol) Get() (bool, error) {
	return t.Table(t.a).Get(t.a)
}

func (t *TablerSymbol) Export() (i interface{}, err error) {
	i = []structure.Symbol{}
	err = t.Table(t.a).Find(&i)

	return
}


