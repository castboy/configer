package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerSource struct {
	bean structure.Xormor
	*xorm.Engine
}

var tablerSource *TablerSource

func NewTablerSource(bean *structure.Source) *TablerSource {
	return &TablerSource{
		bean,
		xEngine,
	}
}

// implement Tableor
func (t *TablerSource) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *TablerSource) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *TablerSource) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *TablerSource) Get() (bool, error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *TablerSource) Export() (i interface{}, err error) {
	i = []structure.Source{}
	err = t.Table(t.bean).Find(&i)

	return
}

func (t *TablerSource) Where() (cond string) {
	return
}


