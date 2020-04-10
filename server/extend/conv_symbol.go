package extend

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type ConvSymboler struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker check.Checkor
}

func NewConvSymboler(tabler *structure.Symbol, bean *structure.ConvSymbol) *ConvSymboler {
	return &ConvSymboler{
		mysql.NewTablerSymbol(tabler),
		cache.NewCacherConvSymbol(bean),
		check.NewCheckerConvSymbol(bean),
	}
}

func (a *ConvSymboler) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *ConvSymboler) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *ConvSymboler) GetChecker() check.Checkor {
	return a.checker
}
