package server

import (
	"configer/server/check"
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
)

// any structure can call functions under this package, if it implement interface below.

type configor interface {
	GetTabler() mysql.TableOperator
	GetCacher() cache.CacheOperator
	GetChecker() check.Checkor
}

