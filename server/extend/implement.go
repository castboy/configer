package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Extender struct {
	tabler  mysql.ExtendOperator
	cacher  cache.BaseOperator
	checker structure.Checkor
}

func NewConvSymboler(tabler *structure.Source, bean *structure.ConvSymbol) *Extender {
	return &Extender{
		mysql.NewTablerSource(tabler),
		cache.NewCacherConvSymbol(bean),
		bean,
	}
}

func NewFullSymbolNamer(tabler *structure.Symbol, bean *structure.FullSymbolName) *Extender {
	return &Extender{
		mysql.NewTablerSymbol(tabler),
		cache.NewCacherFullSymbolName(bean),
		bean,
	}
}

func NewHolidayCalcer(symb *structure.Symbol, bean *structure.HolidayCalc) *Extender {
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
