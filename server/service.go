package server

import (
	"configer/server/check"
	"configer/server/repository"
)

// any structure can call functions under this package, if it implement interface below.

type TableOperator interface {
	repository.BaseOperator
}

type CacheOperator interface {
	repository.BaseOperator

	Cache(i interface{})
}

type configor interface {
	GetTabler() TableOperator
	GetCacher() CacheOperator
	GetChecker() check.Checkor
}

