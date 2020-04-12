package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
)

type TablerSymbol struct {
	*Tabler
}

type Tabler struct {
	bean structure.Xormor
	*xorm.Engine
}

func NewTablerSymbol(bean *structure.Symbol) *TablerSymbol {
	return &TablerSymbol{
		&Tabler{
			bean,
			xEngine,
		},
	}
}

func (t *Tabler) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *Tabler) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *Tabler) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *Tabler) Get() (bool, error) {
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


