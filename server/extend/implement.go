package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
	"configer/server/structure/indexName"
	"configer/server/structure/indexNameID"
)

type Extender struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker structure.Checkor
}

func NewConvSymboler(tabler *indexNameID.Source, bean *indexName.ConvSymbol) *Extender {
	return &Extender{
		mysql.NewTablerSource(tabler),
		cache.NewCacherConvSymbol(bean),
		bean,
	}
}

func NewFullSymbolNamer(tabler *indexNameID.Symbol, bean *structure.FullSymbolName) *Extender {
	return &Extender{
		mysql.NewTablerSymbol(tabler),
		cache.NewCacherFullSymbolName(bean),
		bean,
	}
}

func NewHolidayCalcer(symb *indexNameID.Symbol, bean *structure.HolidayCalc) *Extender {
	return &Extender{
		mysql.NewTablerSymbol(symb),
		cache.NewCacherHolidayCalc(bean),
		bean,
	}
}

func (a *Extender) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *Extender) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *Extender) GetChecker() structure.Checkor {
	return a.checker
}
