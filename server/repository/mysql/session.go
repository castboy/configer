package mysql

import (
	"configer/server/structure"
	"fmt"
	"github.com/go-xorm/xorm"
)

type TablerSession struct {
	bean structure.Xormor
	*xorm.Engine
}

func NewTablerSession(bean *structure.Session) *TablerSession {
	return &TablerSession{
		bean,
		xEngine,
	}
}

// implement Tableor
func (t *TablerSession) Insert() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (t *TablerSession) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (t *TablerSession) Update() (int64, error) {
	return t.Table(t.bean).Update(t.bean)
}

func (t *TablerSession) Get() (exist bool, err error) {
	return t.Table(t.bean).Get(t.bean)
}

func (t *TablerSession) Export() (interface{}, error) {
	i := []structure.Session{}
	err := t.Table(t.bean).Find(&i)

	return i, err
}

func (t *TablerSession) Where() (cond string) {
	return
}



