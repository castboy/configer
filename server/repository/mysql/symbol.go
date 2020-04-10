package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerSymbol struct {
	bean structure.Xormor
	*xorm.Engine
}

var tablerSymbol *TablerSymbol

func NewTablerSymbol(bean *structure.Symbol) *TablerSymbol {
	return &TablerSymbol{
		bean,
		xEngine,
	}
}

// implement Tableor
func (t *TablerSymbol) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *TablerSymbol) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *TablerSymbol) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *TablerSymbol) Get() (bool, error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *TablerSymbol) Export() (i interface{}, err error) {
	i = []structure.Symbol{}
	err = t.Table(t.bean).Find(&i)

	return
}

func (t *TablerSymbol) Where() (cond string) {
	return
}



