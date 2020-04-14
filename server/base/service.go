package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type BaseOperator interface {
	GetTabler() mysql.BaseOperator
	GetCacher() cache.BaseOperator
	GetChecker() structure.Checkor
}
