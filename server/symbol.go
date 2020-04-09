package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type Symboler struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker check.Checkor
}

func NewSymboler(a *structure.Symbol) *Symboler {
	return &Symboler{
		cache.NewCacherSymbol(a),
		mysql.NewTablerSymbol(a),
		check.NewCheckerSymbol(a),
	}
}

func (a *Symboler) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *Symboler) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *Symboler) GetChecker() check.Checkor {
	return a.checker
}

