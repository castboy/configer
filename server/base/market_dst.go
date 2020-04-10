package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type MarketDSTer struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker structure.Checkor
}

func NewMarketDSTer(bean *structure.MarketDST) *MarketDSTer {
	return &MarketDSTer{
		cache.NewCacherMarketDST(bean),
		mysql.NewTablerMarketDST(bean),
		bean,
	}
}

func (a *MarketDSTer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *MarketDSTer) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *MarketDSTer) GetChecker() structure.Checkor {
	return a.checker
}