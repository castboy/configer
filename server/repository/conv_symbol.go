package repository

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/structure"
)

type ConvSymboler struct {
	cacher  cache.BaseOperator
	checker check.Checkor
}

func NewConvSymboler(bean *structure.ConvSymbol) *ConvSymboler {
	return &ConvSymboler{
		cache.NewCacherConvSymbol(bean),
		check.NewCheckerConvSymbol(bean),
	}
}

func (a *ConvSymboler) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *ConvSymboler) GetChecker() check.Checkor {
	return a.checker
}