package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Holidayer struct {
	cacher  cache.CacheOperator
	tabler  mysql.TableOperator
	checker check.Checkor
}

func NewHolidayer(bean *structure.Holiday) *Holidayer {
	return &Holidayer{
		cache.NewCacherHoliday(bean),
		mysql.NewTablerHoliday(bean),
		check.NewCheckerHoliday(bean),
	}
}

func (a *Holidayer) GetCacher() cache.CacheOperator {
	return a.cacher
}

func (a *Holidayer) GetTabler() mysql.TableOperator {
	return a.tabler
}

func (a *Holidayer) GetChecker() check.Checkor {
	return a.checker
}