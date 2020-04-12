package cache

import (
	"configer/server/structure"
)

type CacherSymbol struct {
	bean  *structure.Symbol
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
func (c *symbolCache) insert(symbol *structure.Symbol) {
	c.Lock()
	defer c.Unlock()

	c.ID2Name[symbol.ID] = symbol.Symbol
	c.name2ID[symbol.Symbol] = symbol.ID
	c.info[symbol.ID] = symbol
}

func (c *symbolCache) delete(symbol *structure.Symbol) {
	c.Lock()
	defer c.Unlock()

	if symbol.Symbol == "" {
		symbol.Symbol = c.ID2Name[symbol.ID]
	} else if symbol.ID == 0 {
		symbol.ID = c.name2ID[symbol.Symbol]
	}

	delete(c.name2ID, symbol.Symbol)
	delete(c.ID2Name, symbol.ID)
	delete(c.info, symbol.ID)
}

func (c *symbolCache) update(symbol *structure.Symbol) {
	c.Lock()
	defer c.Unlock()

	if symbol.Symbol == "" {
		symbol.Symbol = c.ID2Name[symbol.ID]
	} else if symbol.ID == 0 {
		symbol.ID = c.name2ID[symbol.Symbol]
	}

	// update fields obviously.
	c.info[symbol.ID].SourceID = symbol.SourceID
	c.info[symbol.ID].Leverage = symbol.Leverage
	c.info[symbol.ID].SecurityID = symbol.SecurityID
	c.info[symbol.ID].MarginInitial = symbol.MarginInitial
	c.info[symbol.ID].MarginDivider = symbol.MarginDivider
	c.info[symbol.ID].Percentage = symbol.Percentage
	c.info[symbol.ID].Status = symbol.Status
}

func (c *symbolCache) get(symbol *structure.Symbol) {
	c.RLock()
	defer c.RUnlock()

	ID := c.name2ID[symbol.Symbol]

	symbol = c.info[ID]
}

func (c *symbolCache) export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}