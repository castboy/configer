package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Sessioner struct {
	cacher  cache.CacheOperator
	tabler  mysql.TableOperator
	checker check.Checkor
}

func NewSessioner(bean *structure.Session) *Sessioner {
	return &Sessioner{
		cache.NewCacherSession(bean),
		mysql.NewTablerSession(bean),
		check.NewCheckerSession(bean),
	}
}

func (a *Sessioner) GetCacher() cache.CacheOperator {
	return a.cacher
}

func (a *Sessioner) GetTabler() mysql.TableOperator {
	return a.tabler
}

func (a *Sessioner) GetChecker() check.Checkor {
	return a.checker
}