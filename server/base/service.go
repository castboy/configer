package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type operator interface {
	GetTabler() mysql.BaseOperator
	GetCacher() cache.BaseOperator
	GetChecker() structure.Helpor
}
