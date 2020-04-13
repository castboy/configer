package cache3

import (
	"configer/server/structure"
	"sync"
)

type BaseCache struct {
	info map[int]structure.Cacheor3
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		info: map[int]structure.Cacheor3{},
	}
}


// cache
func (c *BaseCache) Insert(bean structure.Cacheor3) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *BaseCache) Delete(bean structure.Cacheor3) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	delete(c.info, ID)
}

func (c *BaseCache) Update(bean structure.Cacheor3) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *BaseCache) Get(bean structure.Cacheor3) (res structure.Cacheor3, exist bool) {
	c.RLock()
	defer c.RUnlock()

	ID := bean.GetID()

	res, exist =  c.info[ID]
	return
}

func (c *BaseCache) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
