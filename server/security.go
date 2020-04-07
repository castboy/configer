package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Securityer struct {
	cacher  cache.CacheOperator
	tabler  mysql.TableOperator
	checker check.Checkor
}

func NewSecurityer(bean *structure.Security) *Securityer {
	return &Securityer{
		cache.NewCacherSecurity(bean),
		mysql.NewTablerSecurity(bean),
		check.NewCheckerSecurity(bean),
	}
}

func (a *Securityer) GetCacher() cache.CacheOperator {
	return a.cacher
}

func (a *Securityer) GetTabler() mysql.TableOperator {
	return a.tabler
}

func (a *Securityer) GetChecker() check.Checkor {
	return a.checker
}