package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherSecurity struct {
	bean  *structure.Security
	cache *securityCache
}

var cacherSecurity *CacherSecurity

func NewCacherSecurity(bean *structure.Security) *CacherSecurity {
	if cacherSecurity == nil {
		cacherSecurity = &CacherSecurity{
			bean,
			secCache,
		}
	}

	return cacherSecurity
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
	return
}

func (c *CacherSecurity) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
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