package cache

import (
	"configer/server/structure"
)

type CacherSymbol struct {
	bean  structure.Cacheor
	cache *symbolCache
}

func NewCacherSymbol(bean *structure.Symbol) *CacherSymbol {
	return &CacherSymbol{
		bean,
		symbCache,
	}
}

// implement Cacheor
func (c *CacherSymbol) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherSymbol) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *CacherSymbol) Update() (num int64, err error) {
	c.cache.update(c.bean)
	return
}

func (c *CacherSymbol) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherSymbol) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *CacherSymbol) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		c.cache.insert(&sb[i])
	}
}

// cache
func (c *symbolCache) insert(bean structure.Cacheor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	name := bean.GetName()

	c.ID2Name[ID] = name
	c.name2ID[name] = ID
	c.info[name] = bean
}

func (c *symbolCache) delete(bean structure.Cacheor) {
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

func (c *symbolCache) update(bean structure.Cacheor) {
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

func (c *symbolCache) get(bean structure.Cacheor) structure.Cacheor {
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

func (c *symbolCache) export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}