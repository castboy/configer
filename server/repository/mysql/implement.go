package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
)

type Tabler struct {
	bean structure.Xormor
	*xorm.Engine
}

type TablerSymbol struct {
	*Tabler
}

type TablerHoliday struct {
	*Tabler
}

type TablerSource struct {
	*Tabler
}

type TablerSession struct {
	*Tabler
}

type TablerSecurity struct {
	*Tabler
}

type TablerMarketDST struct {
	*Tabler
}

func NewTablerSymbol(bean *structure.Symbol) *TablerSymbol {
	return &TablerSymbol{
		&Tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerHoliday(bean *structure.Holiday) *TablerHoliday {
	return &TablerHoliday{
		&Tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSource(bean *structure.Source) *TablerSource {
	return &TablerSource{
		&Tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSession(bean *structure.Session) *TablerSession {
	return &TablerSession{
		&Tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSecurity(bean *structure.Security) *TablerSecurity {
	return &TablerSecurity{
		&Tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerMarketDST(bean *structure.MarketDST) *TablerMarketDST {
	return &TablerMarketDST{
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

func (t *TablerHoliday) Export() (interface{}, error) {
	i := []structure.Holiday{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *TablerSource) Export() (interface{}, error) {
	i := []structure.Source{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *TablerSession) Export() (interface{}, error) {
	i := []structure.Session{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *TablerSecurity) Export() (interface{}, error) {
	i := []structure.Security{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *TablerMarketDST) Export() (interface{}, error) {
	i := []structure.MarketDST{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}