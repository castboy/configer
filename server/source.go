package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Sourcer struct {
	cacher  cache.CacheOperator
	tabler  mysql.TableOperator
	checker check.Checkor
}

func NewSourcer(bean *structure.Source) *Sourcer {
	return &Sourcer{
		cache.NewCacherSource(bean),
		mysql.NewTablerSource(bean),
		check.NewCheckerSource(bean),
	}
}

func (a *Sourcer) GetCacher() cache.CacheOperator {
	return a.cacher
}

func (a *Sourcer) GetTabler() mysql.TableOperator {
	return a.tabler
}

func (a *Sourcer) GetChecker() check.Checkor {
	return a.checker
}
