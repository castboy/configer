package server

import "configer/server/check"

// any structure can call functions under this package, if it implement interface below.

type TableOperator interface {
	baseOperator
}

type CacheOperator interface {
	baseOperator

	Cache(i interface{})
}

type baseOperator interface {
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	Get() (bool, error)
	Export() (i interface{}, err error)
}

type configor interface {
	GetTabler() TableOperator
	GetCacher() CacheOperator
	GetChecker() check.Checkor
}

