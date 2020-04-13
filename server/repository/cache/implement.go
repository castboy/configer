package cache

import (
	"configer/server/structure"
)

type CacherSymbol struct {
	*Cacher
}

type CacherSource struct {
	*Cacher
}

type CacherSecurity struct {
	*Cacher
}

type Cacher struct {
	bean  structure.Cacheor
	cache *baseCache
}

func NewCacherSymbol(bean *structure.Symbol) *CacherSymbol {
	return &CacherSymbol{
		&Cacher{
			bean,
			symbCache,
		},
	}
}

func NewCacherSource(bean *structure.Source) *CacherSource {
	return &CacherSource{
		&Cacher{
			bean,
			srcCache,
		},
	}
}

func NewCacherSecurity(bean *structure.Security) *CacherSecurity {
	return &CacherSecurity{
		&Cacher{
			bean,
			secCache,
		},
	}
}

// implement Cacheor
func (c *Cacher) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *Cacher) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *Cacher) Update() (num int64, err error) {
	c.cache.update(c.bean)
	return
}

func (c *Cacher) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *Cacher) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *CacherSymbol) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		c.cache.insert(&sb[i])
	}
}

func (c *CacherSource) Cache(i interface{}) {
	src := i.([]structure.Source)
	for i := range src {
		c.cache.insert(&src[i])
	}
}

func (c *CacherSecurity) Cache(i interface{}) {
	se := i.([]structure.Security)
	for i := range se {
		c.cache.insert(&se[i])
	}
}

// cache
func (c *baseCache) insert(bean structure.Cacheor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	name := bean.GetName()

	c.ID2Name[ID] = name
	c.name2ID[name] = ID
	c.info[name] = bean
}

func (c *baseCache) delete(bean structure.Cacheor) {
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

func (c *baseCache) update(bean structure.Cacheor) {
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

func (c *baseCache) get(bean structure.Cacheor) structure.Cacheor {
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

func (c *baseCache) export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}