package repository

import (
	"configer/server/check"
	"configer/server/repository/cache"
)

type BaseOperator interface {
	GetCacher() cache.BaseOperator
	GetChecker() check.Checkor
}