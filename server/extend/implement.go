package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
	"configer/server/structure/indexName"
	"configer/server/structure/indexNameID"
)

type extender struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker structure.Checkor
}

func NewConvSymboler(tabler *indexNameID.Source, bean *indexName.ConvSymbol) *extender {
	return &extender{
		mysql.NewTablerSource(tabler),
		cache.NewCacherConvSymbol(bean),
		bean,
	}
}

func NewFullSymbolNamer(tabler *indexNameID.Symbol, bean *structure.FullSymbolName) *extender {
	return &extender{
		mysql.NewTablerSymbol(tabler),
		cache.NewCacherFullSymbolName(bean),
		bean,
	}
}

func NewHolidayCalcer(symb *indexNameID.Symbol, bean *structure.HolidayCalc) *extender {
	return &extender{
		mysql.NewTablerSymbol(symb),
		cache.NewCacherHolidayCalc(bean),
		bean,
	}
}

func (a *extender) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *extender) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *extender) GetChecker() structure.Checkor {
	return a.checker
}
