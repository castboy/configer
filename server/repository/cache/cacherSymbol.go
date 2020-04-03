package cache

import "configer/server/structure"

type CacherSymbol struct {
	a *structure.Symbol
	b *symbolCache
}

var cacherSymbol *CacherSymbol

func NewCacherSymbol(a *structure.Symbol) *CacherSymbol {
	if cacherSymbol == nil {
		cacherSymbol = &CacherSymbol{
			a,
			symbCache,
		}
	}

	return cacherSymbol
}

// implement Cacheor
func (c *CacherSymbol) Insert() (num int64, err error) {
	c.b.insert(c.a)
	return
}

func (c *CacherSymbol) Delete() (num int64, err error) {
	c.b.delete(c.a)
	return
}

func (c *CacherSymbol) Update() (num int64, err error) {
	c.b.update(c.a)
	return
}

func (c *CacherSymbol) Get() (exist bool, err error) {
	c.b.get(c.a)
	return
}

func (c *CacherSymbol) Export() (i interface{}, err error) {
	return
}

func (c *CacherSymbol) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
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