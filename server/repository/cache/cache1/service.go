package cache1

import "configer/server/structure"

type Cache1 interface {
	Insert(bean structure.Cacheor1)
	Delete(bean structure.Cacheor1)
	Update(bean structure.Cacheor1)
	Get(bean structure.Cacheor1) structure.Cacheor1
	Export() (i interface{}, err error)
}

