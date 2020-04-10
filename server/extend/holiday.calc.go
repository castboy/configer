package extend

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type HolidayCalcer struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker check.Checkor
}

func NewHolidayCalcer(symb *structure.Symbol, bean *structure.HolidayCalc) *HolidayCalcer {
	return &HolidayCalcer{
		mysql.NewTablerSymbol(symb),
		cache.NewCacherHolidayCalc(bean),
		check.NewCheckerHolidayCalc(bean),
	}
}

func (a *HolidayCalcer) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *HolidayCalcer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *HolidayCalcer) GetChecker() check.Checkor {
	return a.checker
}

