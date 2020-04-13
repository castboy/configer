package cache2

import (
	"configer/server/structure"
	"sync"
)

type BaseCache struct {
	info map[string]structure.Cacheor2
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		info: map[string]structure.Cacheor2{},
	}
}

// cache
func (c *BaseCache) Insert(bean structure.Cacheor2) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *BaseCache) Delete(bean structure.Cacheor2) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	delete(c.info, name)
}

func (c *BaseCache) Update(bean structure.Cacheor2) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *BaseCache) Get(bean structure.Cacheor2) (res structure.Cacheor2, exist bool) {
	c.RLock()
	defer c.RUnlock()

	name := bean.GetName()

	res, exist =  c.info[name]
	return
}

func (c *BaseCache) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}