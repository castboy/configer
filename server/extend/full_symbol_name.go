package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type FullSymbolNamer struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker structure.Checkor
}

func NewFullSymbolNamer(tabler *structure.Symbol, bean *structure.FullSymbolName) *FullSymbolNamer {
	return &FullSymbolNamer{
		mysql.NewTablerSymbol(tabler),
		cache.NewCacherFullSymbolName(bean),
		bean,
	}
}

func (a *FullSymbolNamer) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *FullSymbolNamer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *FullSymbolNamer) GetChecker() structure.Checkor {
	return a.checker
}