package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherConvSymbol struct {
	bean *structure.ConvSymbol
	cache *convSymbolCache
}

var cacherConvSymbol *CacherConvSymbol

func NewCacherConvSymbol(bean *structure.ConvSymbol) *CacherConvSymbol {
	if cacherConvSymbol == nil {
		cacherConvSymbol = &CacherConvSymbol{
			bean: bean,
			cache: csCache[bean.ConvType],
		}
	}

	return cacherConvSymbol
}

// implement Cacheor
func (c *CacherConvSymbol) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherConvSymbol) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherConvSymbol) Update() (num int64, err error) {
	return
}

func (c *CacherConvSymbol) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherConvSymbol) Export() (i interface{}, err error) {
	return
}

func (c *CacherConvSymbol) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *convSymbolCache) insert(cs *structure.ConvSymbol) {
	c.Lock()
	defer c.Unlock()

	c.info[cs.SourceName] = cs
}

func (c *convSymbolCache) get(cs *structure.ConvSymbol) {
	c.RLock()
	defer c.RUnlock()

	cs = c.info[cs.SourceName]
}