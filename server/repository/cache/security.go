package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherSecurity struct {
	bean  *structure.Security
	cache *securityCache
}

func NewCacherSecurity(bean *structure.Security) *CacherSecurity {
	return &CacherSecurity{
		bean,
		secCache,
	}
}

// implement Cacheor
func (c *CacherSecurity) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherSecurity) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherSecurity) Update() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherSecurity) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherSecurity) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *CacherSecurity) Cache(i interface{}) {
	se := i.([]structure.Security)
	for i := range se {
		c.cache.insert(&se[i])
	}
}

// cache
func (c *securityCache) insert(security *structure.Security) {
	c.Lock()
	defer c.Unlock()

	c.ID2Name[security.ID] = security.SecurityName
	c.name2ID[security.SecurityName] = security.ID
	c.info[security.SecurityName] = security
}

func (c *securityCache) get(security *structure.Security) {
	c.RLock()
	defer c.RUnlock()

	if security.SecurityName == "" {
		security.SecurityName = c.ID2Name[security.ID]
	}

	security = c.info[security.SecurityName]
}

func (c *securityCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}