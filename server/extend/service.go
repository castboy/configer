package extend

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type operator interface {
	GetTabler() mysql.ExtendOperator
	GetCacher() cache.BaseOperator
	GetHelper() structure.Helpor
}
