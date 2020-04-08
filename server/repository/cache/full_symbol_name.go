package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherFullSymbolName struct {
	bean *structure.FullSymbolName
	cache *fullSymbolNameCache
}


var cacherFullSymbolName *CacherFullSymbolName

func NewCacherFullSymbolName(bean *structure.FullSymbolName) *CacherFullSymbolName {
	if cacherFullSymbolName == nil {
		cacherFullSymbolName = &CacherFullSymbolName{
			bean,
			fsnCache,
		}
	}

	return cacherFullSymbolName
}

// implement Cacheor
func (c *CacherFullSymbolName) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherFullSymbolName) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *CacherFullSymbolName) Update() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherFullSymbolName) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherFullSymbolName) Export() (i interface{}, err error) {
	return
}

func (c *CacherFullSymbolName) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *fullSymbolNameCache) insert(fsn *structure.FullSymbolName) {
	c.Lock()
	defer c.Unlock()

	c.info[fsn.Sl] = fsn
}

func (c *fullSymbolNameCache) delete(fsn *structure.FullSymbolName) {
	c.Lock()
	defer c.Unlock()

	delete(c.info, fsn.Sl)
}

func (c *fullSymbolNameCache) get(fsn *structure.FullSymbolName) {
	c.RLock()
	defer c.RUnlock()

	fsn = c.info[fsn.Sl]
}