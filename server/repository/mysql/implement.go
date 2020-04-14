package mysql

import (
	"configer/server/structure"
	"configer/server/structure/indexID"
	"configer/server/structure/indexNameID"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
)

type tabler struct {
	bean structure.Xormor
	*xorm.Engine
}

type tablerSymbol struct {
	*tabler
}

type tablerHoliday struct {
	*tabler
}

type tablerSource struct {
	*tabler
}

type tablerSession struct {
	*tabler
}

type tablerSecurity struct {
	*tabler
}

type tablerMarketDST struct {
	*tabler
}

func NewTablerSymbol(bean *indexNameID.Symbol) *tablerSymbol {
	return &tablerSymbol{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerHoliday(bean *indexID.Holiday) *tablerHoliday {
	return &tablerHoliday{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSource(bean *indexNameID.Source) *tablerSource {
	return &tablerSource{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSession(bean *structure.Session) *tablerSession {
	return &tablerSession{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSecurity(bean *indexNameID.Security) *tablerSecurity {
	return &tablerSecurity{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerMarketDST(bean *structure.MarketDST) *tablerMarketDST {
	return &tablerMarketDST{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func (t *tabler) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *tabler) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *tabler) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *tabler) Get() (bool, error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *tablerSymbol) Export() (interface{}, error) {
	i := []indexNameID.Symbol{}
	err := t.Table(t.bean).Find(&i)

	for j := range i {
		i[j].Leverage = int32(i[j].MarginDivider.Mul(decimal.NewFromFloat(100)).IntPart())
	}

	return i, err
}

func (t *tablerHoliday) Export() (interface{}, error) {
	i := []indexID.Holiday{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerSource) Export() (interface{}, error) {
	i := []indexNameID.Source{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerSession) Export() (interface{}, error) {
	i := []structure.Session{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerSecurity) Export() (interface{}, error) {
	i := []indexNameID.Security{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerMarketDST) Export() (interface{}, error) {
	i := []structure.MarketDST{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}
