package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherMarketDST struct {
	bean  *structure.MarketDST
	cache *marketDSTCache
}

var cacherMarketDST *CacherMarketDST

func NewCacherMarketDST(bean *structure.MarketDST) *CacherMarketDST {
	if cacherMarketDST == nil {
		cacherMarketDST = &CacherMarketDST{
			bean,
			mdCache,
		}
	}

	return cacherMarketDST
}

// implement Cacheor
func (c *CacherMarketDST) Insert() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherMarketDST) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherMarketDST) Update() (num int64, err error) {
	c.cache.set(c.bean)
	return
}

func (c *CacherMarketDST) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherMarketDST) Export() (i interface{}, err error) {
	return
}

func (c *CacherMarketDST) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *marketDSTCache) set(marketDST *structure.MarketDST) {
	c.Lock()
	defer c.Unlock()

	c.info[marketDST.MarketOwnerType] = marketDST
}

func (c *marketDSTCache) get(marketDST *structure.MarketDST) {
	c.RLock()
	defer c.RUnlock()

	marketDST = c.info[marketDST.MarketOwnerType]
}