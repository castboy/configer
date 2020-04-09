package mysql

import (
	"configer/server/structure"
	"github.com/go-xorm/xorm"
)

type TablerHoliday struct {
	bean *structure.Holiday
	*xorm.Engine
}

var tablerHoliday *TablerHoliday

func NewTablerHoliday(bean *structure.Holiday) *TablerHoliday {
	return &TablerHoliday{
		bean,
		xEngine,
	}
}

// implement Tableor
func (t *TablerHoliday) Insert() (int64, error) {
	return t.Table(t.bean).Insert(t.bean)
}

func (t *TablerHoliday) Delete() (int64, error) {
	return t.Table(t.bean).Delete(t.bean)
}

func (t *TablerHoliday) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *TablerHoliday) Get() (bool, error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *TablerHoliday) Export() (i interface{}, err error) {
	i = []structure.Holiday{}
	err = t.Table(t.bean).Find(&i)

	return
}


