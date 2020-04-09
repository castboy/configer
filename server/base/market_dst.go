package base

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type MarketDSTer struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker check.Checkor
}

func NewMarketDSTer(bean *structure.MarketDST) *MarketDSTer {
	return &MarketDSTer{
		cache.NewCacherMarketDST(bean),
		mysql.NewTablerMarketDST(bean),
		check.NewCheckerMarketDST(bean),
	}
}

func (a *MarketDSTer) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *MarketDSTer) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *MarketDSTer) GetChecker() check.Checkor {
	return a.checker
}