package cache

import (
	"configer/server/structure"
	"configer/server/utils"
	"fmt"
)

type CacherConvSymbol struct {
	bean *structure.ConvSymbol
	cache *convSymbolCache
}

func NewCacherConvSymbol(bean *structure.ConvSymbol) *CacherConvSymbol {
	return &CacherConvSymbol{
		bean: bean,
		cache: csCache[bean.ConvType],
	}
}

// implement Cacheor1
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
	return c.cache.export()
}

func (c *CacherConvSymbol) Cache(i interface{}) {
	src := i.([]structure.Source)
	for i := range src {
		c.bean.ConvInfo = utils.BuildConvInfo(src[i].Source, src)
		c.bean.SourceName = src[i].Source

		c.cache.insert(c.bean)
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

func (c *convSymbolCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}