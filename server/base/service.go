package base

import (
	"configer/server/structure"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
)

type BaseOperator interface {
	GetTabler() mysql.BaseOperator
	GetCacher() cache.BaseOperator
	GetChecker() structure.Checkor
}

