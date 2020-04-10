package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Holidayer struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker structure.Checkor
}

func NewHolidayer(bean *structure.Holiday) *Holidayer {
	return &Holidayer{
		cache.NewCacherHoliday(bean),
		mysql.NewTablerHoliday(bean),
		bean,
	}
}

func (a *Holidayer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *Holidayer) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *Holidayer) GetChecker() structure.Checkor {
	return a.checker
}