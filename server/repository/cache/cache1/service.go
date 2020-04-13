package cache1

import "configer/server/structure"

type Cache1 interface {
	Insert(bean structure.Cacheor)
	Delete(bean structure.Cacheor)
	Update(bean structure.Cacheor)
	Get(bean structure.Cacheor) structure.Cacheor
	Export() (i interface{}, err error)
}

