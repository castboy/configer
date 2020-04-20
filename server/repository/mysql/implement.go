package mysql

import (
	"configer/server/structure"
	"configer/server/utils"
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

func NewTablerSymbol(bean *structure.Symbol) *tablerSymbol {
	return &tablerSymbol{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerHoliday(bean *structure.Holiday) *tablerHoliday {
	return &tablerHoliday{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func NewTablerSource(bean *structure.Source) *tablerSource {
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

func NewTablerSecurity(bean *structure.Security) *tablerSecurity {
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
	return t.Table(t.bean).Where(t.bean.DeleteCondition()).NoAutoCondition(true).Delete(t.bean)
}

func (t *tabler) Update() (int64, error) {
	return t.Table(t.bean).Where(t.bean.UpdateCondition()).Update(t.bean)
}

func (t *tabler) Get() (bool, error) {
	return t.Table(t.bean).Where(t.bean.GetCondition()).Get(t.bean)
}

func (t *tablerSymbol) Export() (interface{}, error) {
	i := []structure.Symbol{}
	err := t.Table(t.bean).Find(&i)

	for j := range i {
		i[j].Leverage = int32(i[j].MarginDivider.Mul(decimal.NewFromFloat(100)).IntPart())
	}

	return i, err
}

func (t *tablerHoliday) Export() (interface{}, error) {
	i := []structure.Holiday{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerSource) Export() (interface{}, error) {
	i := []structure.Source{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerSession) Update() (num int64, err error) {
	ses := t.bean.(*structure.Session)
	ses.Session = utils.OrderAndFill(ses.Session)

	return t.Table(t.bean).Where(t.bean.UpdateCondition()).Update(ses)
}

func (t *tablerSession) Export() (interface{}, error) {
	i := []structure.Session{}
	err := t.Table(t.bean).Where(t.bean.ExportCondition()).Find(&i)

	return i, err
}

func (t *tablerSecurity) Export() (interface{}, error) {
	i := []structure.Security{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *tablerMarketDST) Export() (interface{}, error) {
	i := []structure.MarketDST{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}
