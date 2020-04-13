package cache

import (
	"configer/server/structure"
	"configer/server/utils"
	"fmt"
)

type CacherFullSymbolName struct {
	bean *structure.FullSymbolName
	cache *fullSymbolNameCache
}

func NewCacherFullSymbolName(bean *structure.FullSymbolName) *CacherFullSymbolName {
	return &CacherFullSymbolName{
		bean,
		fsnCache,
	}
}

// implement Cacheor1
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
	return c.cache.export()
}

func (c *CacherFullSymbolName) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		fsn := &structure.FullSymbolName{}
		fsn.Sl.Symbol = utils.GetRequestSymbol(sb[i].Symbol)
		fsn.Sl.Leverage = sb[i].Leverage
		fsn.FullName = sb[i].Symbol

		c.cache.insert(fsn)
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

func (c *fullSymbolNameCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}