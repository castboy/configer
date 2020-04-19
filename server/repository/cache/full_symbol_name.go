package cache

import (
	"configer/server/structure"
	"configer/server/utils"
)

type cacherFullSymbolName struct {
	bean  *structure.FullSymbolName
	cache *fullSymbolNameCache
}

func NewCacherFullSymbolName(bean *structure.FullSymbolName) *cacherFullSymbolName {
	return &cacherFullSymbolName{
		bean,
		fsnCache,
	}
}

func (c *cacherFullSymbolName) Insert() {
	c.cache.insert(c.bean)
}

func (c *cacherFullSymbolName) Delete() {
	c.cache.delete(c.bean)
}

func (c *cacherFullSymbolName) Update() {
}

func (c *cacherFullSymbolName) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *cacherFullSymbolName) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *cacherFullSymbolName) Cache(i interface{}) {
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
