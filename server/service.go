package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
)

type Configor interface {
	GetTabler() mysql.TableOperator
	GetCacher() cache.CacheOperator
	GetChecker() check.Checkor
}

type ExtendOperator interface {
	GetTabler() mysql.ExtendOperator
	GetCacher() cache.ExtendOperator
	GetChecker() check.Checkor
}

