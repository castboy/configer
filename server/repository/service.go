package repository

import (
	"configer/server/check"
	"configer/server/repository/cache"
)

type Configor interface {
	GetCacher() cache.CacheOperator
	GetChecker() check.Checkor
}