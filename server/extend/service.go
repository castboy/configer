package extend

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
)

type ExtendOperator interface {
	GetTabler() mysql.ExtendOperator
	GetCacher() cache.BaseOperator
	GetChecker() check.Checkor
}

