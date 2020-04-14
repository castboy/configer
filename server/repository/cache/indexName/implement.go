package indexName

import (
	"configer/server/structure"
	"sync"
)

type BaseCache struct {
	info map[string]structure.Nameor
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		info: map[string]structure.Nameor{},
	}
}

// cache
func (c *BaseCache) Insert(bean structure.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *BaseCache) Delete(bean structure.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	delete(c.info, name)
}

func (c *BaseCache) Update(bean structure.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *BaseCache) Get(bean structure.Nameor) (res structure.Nameor, exist bool) {
	c.RLock()
	defer c.RUnlock()

	name := bean.GetName()

	res, exist = c.info[name]
	return
}

func (c *BaseCache) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
