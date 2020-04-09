package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
)

type BaseOperator interface {
	GetTabler() mysql.BaseOperator
	GetCacher() cache.BaseOperator
	GetChecker() check.Checkor
}

type ExtendOperator interface {
	GetTabler() mysql.ExtendOperator
	GetCacher() cache.ExtendOperator
	GetChecker() check.Checkor
}

