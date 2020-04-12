package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type ConvSymboler struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker structure.Checkor
}

func NewConvSymboler(tabler *structure.Source, bean *structure.ConvSymbol) *ConvSymboler {
	return &ConvSymboler{
		mysql.NewTablerSource(tabler),
		cache.NewCacherConvSymbol(bean),
		bean,
	}
}

func (a *ConvSymboler) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *ConvSymboler) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *ConvSymboler) GetChecker() structure.Checkor {
	return a.checker
}
