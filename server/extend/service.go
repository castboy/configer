package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type ExtendOperator interface {
	GetTabler() mysql.ExtendOperator
	GetCacher() cache.BaseOperator
	GetChecker() structure.Checkor
}
