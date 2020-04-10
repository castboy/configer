package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Sessioner struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker structure.Checkor
}

func NewSessioner(bean *structure.Session) *Sessioner {
	return &Sessioner{
		cache.NewCacherSession(bean),
		mysql.NewTablerSession(bean),
		bean,
	}
}

func (a *Sessioner) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *Sessioner) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *Sessioner) GetChecker() structure.Checkor {
	return a.checker
}