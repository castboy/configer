package cache2

import "configer/server/structure"

type Cache2 interface {
	Insert(bean structure.Cacheor2)
	Delete(bean structure.Cacheor2)
	Update(bean structure.Cacheor2)
	Get(bean structure.Cacheor2) (structure.Cacheor2, bool)
	Export() (i interface{}, err error)
}

