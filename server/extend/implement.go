package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type extender struct {
	tabler mysql.ExtendOperator
	cacher cache.BaseOperator
	helper structure.Helpor
}

func NewConvSymboler(bean *structure.ConvSymbol) *extender {
	return &extender{
		mysql.NewTablerSource(&structure.Source{}),
		cache.NewCacherConvSymbol(bean),
		bean,
	}
}

func NewFullSymbolNamer(bean *structure.FullSymbolName) *extender {
	return &extender{
		mysql.NewTablerSymbol(&structure.Symbol{}),
		cache.NewCacherFullSymbolName(bean),
		bean,
	}
}

func (a *extender) GetTabler() mysql.ExtendOperator {
	return a.tabler
}

func (a *extender) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *extender) GetHelper() structure.Helpor {
	return a.helper
}
