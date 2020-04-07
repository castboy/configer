package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type MarketDSTer struct {
	cacher  cache.CacheOperator
	tabler  mysql.TableOperator
	checker check.Checkor
}

func NewMarketDSTer(bean *structure.MarketDST) *MarketDSTer {
	return &MarketDSTer{
		cache.NewCacherMarketDST(bean),
		mysql.NewTablerMarketDST(bean),
		check.NewCheckerMarketDST(bean),
	}
}

func (a *MarketDSTer) GetCacher() cache.CacheOperator {
	return a.cacher
}

func (a *MarketDSTer) GetTabler() mysql.TableOperator {
	return a.tabler
}

func (a *MarketDSTer) GetChecker() check.Checkor {
	return a.checker
}