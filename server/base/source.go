package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Sourcer struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker structure.Checkor
}

func NewSourcer(bean *structure.Source) *Sourcer {
	return &Sourcer{
		cache.NewCacherSource(bean),
		mysql.NewTablerSource(bean),
		bean,
	}
}

func (a *Sourcer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *Sourcer) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *Sourcer) GetChecker() structure.Checkor {
	return a.checker
}
