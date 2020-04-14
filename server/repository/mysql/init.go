package mysql

import (
	"configer/server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var xEngine *xorm.Engine

func init() {
	var err error
	xEngine, err = xorm.NewEngine("mysql", "root:wang1234@/trading_system?charset=utf8")
	if err != nil {
		panic(err)
	}

	xEngine.ShowSQL(true)

	xEngine.SetTableMapper(utils.SetObjTables(utils.NewObjTable("MarketDST", "market_dst")))
}
