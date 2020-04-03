package cache

import "configer/server/repository"

type CacheOperator interface {
	repository.BaseOperator

	Cache(i interface{})
}

