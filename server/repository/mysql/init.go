package mysql

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var xEngine *xorm.Engine

func init() {
	var err error
	xEngine, err = xorm.NewEngine("mysql", "root:wang1234@/symbol_test?charset=utf8")
	if err != nil {
		panic(err)
	}
}
