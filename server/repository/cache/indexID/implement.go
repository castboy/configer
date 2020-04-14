package indexID

import (
	"configer/server/structure/indexID"
	"sync"
)

type BaseCache struct {
	info map[int]indexID.IDor
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		info: map[int]indexID.IDor{},
	}
}

// cache
func (c *BaseCache) Insert(bean indexID.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *BaseCache) Delete(bean indexID.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	delete(c.info, ID)
}

func (c *BaseCache) Update(bean indexID.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *BaseCache) Get(bean indexID.IDor) (res indexID.IDor, exist bool) {
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
