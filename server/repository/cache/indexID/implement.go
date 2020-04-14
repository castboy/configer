package indexID

import (
	"configer/server/structure"
	"sync"
)

type BaseCache struct {
	info map[int]structure.IDor
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		info: map[int]structure.IDor{},
	}
}

// cache
func (c *BaseCache) Insert(bean structure.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *BaseCache) Delete(bean structure.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	delete(c.info, ID)
}

func (c *BaseCache) Update(bean structure.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *BaseCache) Get(bean structure.IDor) (res structure.IDor, exist bool) {
	c.RLock()
	defer c.RUnlock()

	ID := bean.GetID()

	res, exist = c.info[ID]
	return
}

func (c *BaseCache) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
