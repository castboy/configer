package indexName

import (
	"configer/server/structure/indexName"
	"sync"
)

type BaseCache struct {
	info map[string]indexName.Nameor
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		info: map[string]indexName.Nameor{},
	}
}

// cache
func (c *BaseCache) Insert(bean indexName.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *BaseCache) Delete(bean indexName.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	delete(c.info, name)
}

func (c *BaseCache) Update(bean indexName.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *BaseCache) Get(bean indexName.Nameor) (res indexName.Nameor, exist bool) {
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
