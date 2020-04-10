package extend

import (
	"configer/server/structure"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
)

type ExtendOperator interface {
	GetTabler() mysql.ExtendOperator
	GetCacher() cache.BaseOperator
	GetChecker() structure.Checkor
}

