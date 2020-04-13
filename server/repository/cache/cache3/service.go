package cache3

import "configer/server/structure"

type Cache3 interface {
	Insert(bean structure.Cacheor3)
	Delete(bean structure.Cacheor3)
	Update(bean structure.Cacheor3)
	Get(bean structure.Cacheor3) (structure.Cacheor3, bool)
	Export() (i interface{}, err error)
}