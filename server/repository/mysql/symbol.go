package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
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

func (t *TablerSymbol) Export() (interface{}, error) {
	i := []structure.Symbol{}
	err := t.Table(t.bean).Find(&i)

	for j := range i {
		i[j].Leverage = int32(i[j].MarginDivider.Mul(decimal.NewFromFloat(100)).IntPart())
	}

	return i, err
}

func (t *TablerSymbol) Where() (cond string) {
	return
}



