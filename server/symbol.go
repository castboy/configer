package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

// Symboler as an implement of interface defined in this package.

type Symboler struct {
	cacherSymbol cacheOperator
	tablerSymbol tableOperator
	checkerSymbol checkor
}

func NewSymboler(a *structure.Symbol) *Symboler {
	return &Symboler{
		cache.NewCacherSymbol(a),
		mysql.NewTablerSymbol(a),
		check.NewCheckerSymbol(a),
	}
}

func (a *Symboler) GetCacher() cacheOperator {
	return a.cacherSymbol
}

func (a *Symboler) GetTabler() tableOperator {
	return a.tablerSymbol
}

func (a *Symboler) GetChecker() checkor {
	return a.checkerSymbol
}

