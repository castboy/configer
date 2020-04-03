package server

import (
	"configer/server/check"
	"configer/server/repository"
	"configer/server/repository/mysql"
)

// any structure can call functions under this package, if it implement interface below.

type CacheOperator interface {
	repository.BaseOperator

	Cache(i interface{})
}

type configor interface {
	GetTabler() mysql.TableOperator
	GetCacher() CacheOperator
	GetChecker() check.Checkor
}

