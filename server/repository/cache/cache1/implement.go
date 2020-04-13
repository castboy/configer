package cache1

import (
	"configer/server/structure"
	"sync"
)

type BaseCache struct {
	ID2Name map[int]string
	name2ID map[string]int
	info    map[string]structure.Cacheor
	sync.RWMutex
}

func NewBaseCache() *BaseCache {
	return &BaseCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]structure.Cacheor),
	}
}
// cache
func (c *BaseCache) Insert(bean structure.Cacheor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	name := bean.GetName()

	c.ID2Name[ID] = name
	c.name2ID[name] = ID
	c.info[name] = bean
}

func (c *BaseCache) Delete(bean structure.Cacheor) {
	c.Lock()
	defer c.Unlock()

	var ID int
	var name string

	ID = bean.GetID()
	if ID != 0 {
		name = c.ID2Name[ID]
	} else {
		name = bean.GetName()
		ID = c.name2ID[name]
	}

	delete(c.name2ID, name)
	delete(c.ID2Name, ID)
	delete(c.info, name)
}

func (c *BaseCache) Update(bean structure.Cacheor) {
	c.Lock()
	defer c.Unlock()

	var ID int
	var name string

	ID = bean.GetID()
	if ID != 0 {
		name = c.ID2Name[ID]
	} else {
		name = bean.GetName()
		ID = c.name2ID[name]
	}

	c.info[name] = bean
}

func (c *BaseCache) Get(bean structure.Cacheor) structure.Cacheor {
	c.RLock()
	defer c.RUnlock()

	var ID int
	var name string

	ID = bean.GetID()
	if ID != 0 {
		name = c.ID2Name[ID]
	} else {
		name = bean.GetName()
		ID = c.name2ID[name]
	}

	return c.info[name]
}

func (c *BaseCache) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
