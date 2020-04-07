package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Symboler struct {
	cacherSymbol  cache.CacheOperator
	tablerSymbol  mysql.TableOperator
	checkerSymbol check.Checkor
}

func NewSymboler(a *structure.Symbol) *Symboler {
	return &Symboler{
		cache.NewCacherSymbol(a),
		mysql.NewTablerSymbol(a),
		check.NewCheckerSymbol(a),
	}
}

func (a *Symboler) GetCacher() cache.CacheOperator {
	return a.cacherSymbol
}

func (a *Symboler) GetTabler() mysql.TableOperator {
	return a.tablerSymbol
}

func (a *Symboler) GetChecker() check.Checkor {
	return a.checkerSymbol
}

