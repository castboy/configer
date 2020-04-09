package repository

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/structure"
)

type FullSymbolNamer struct {
	cacher  cache.BaseOperator
	checker check.Checkor
}

func NewFullSymbolNamer(bean *structure.FullSymbolName) *FullSymbolNamer {
	return &FullSymbolNamer{
		cache.NewCacherFullSymbolName(bean),
		check.NewCheckerFullSymbolName(bean),
	}
}

func (a *FullSymbolNamer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *FullSymbolNamer) GetChecker() check.Checkor {
	return a.checker
}
