package server

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

// Aer as an implement of interface defined in this package.

type Aer struct {
	cacherA operator
	tablerA operator
}

func NewAer(a *structure.A) *Aer {
	return &Aer{
		cache.NewCacherA(a),
		mysql.NewTablerA(a),
	}
}

func (a *Aer) GetCacher() operator {
	return a.cacherA
}

func (a *Aer) GetTabler() operator {
	return a.tablerA
}

