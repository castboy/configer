package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type FullSymbolNamer struct {
	tabler  mysql.ExtendOperator
	cacher  cache.ExtendOperator
	checker check.Checkor
}

func NewFullSymbolNamer(tabler *structure.Symbol, bean *structure.FullSymbolName) *FullSymbolNamer {
	return &FullSymbolNamer{
		mysql.NewTablerSymbol(tabler),
		cache.NewCacherFullSymbolName(bean),
		check.NewCheckerFullSymbolName(bean),
	}
}

func (a *FullSymbolNamer) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *FullSymbolNamer) GetCacher() cache.ExtendOperator {
	return a.cacher
}

func (a *FullSymbolNamer) GetChecker() check.Checkor {
	return a.checker
}