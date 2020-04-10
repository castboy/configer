package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerSecurity struct {
	bean *structure.Security
	*xorm.Engine
}

var tablerSecurity *TablerSecurity

func NewTablerSecurity(bean *structure.Security) *TablerSecurity {
	return &TablerSecurity{
		bean,
		xEngine,
	}
}

// implement Tableor
func (t *TablerSecurity) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *TablerSecurity) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *TablerSecurity) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *TablerSecurity) Get() (bool, error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *TablerSecurity) Export() (i interface{}, err error) {
	i = []structure.Security{}
	err = t.Table(t.bean).Find(&i)

	return
}

func (t *TablerSecurity) Where() (cond string) {
	return
}

