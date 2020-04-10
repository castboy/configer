package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerMarketDST struct {
	bean structure.Xormor
	*xorm.Engine
}

var tablerMarketDST *TablerMarketDST

func NewTablerMarketDST(bean *structure.MarketDST) *TablerMarketDST {
	return &TablerMarketDST{
		bean,
		xEngine,
	}
}

// implement Tableor
func (t *TablerMarketDST) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *TablerMarketDST) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *TablerMarketDST) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *TablerMarketDST) Get() (bool, error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *TablerMarketDST) Export() (i interface{}, err error) {
	i = []structure.MarketDST{}
	err = t.Table(t.bean).Find(&i)

	return
}

func (t *TablerMarketDST) Where() (cond string) {
	return
}


