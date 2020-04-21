package cache

import (
	"configer/server/structure"
)

type cacherMarketDST struct {
	bean  *structure.MarketDST
	cache *marketDSTCache
}

func NewCacherMarketDST(bean *structure.MarketDST) *cacherMarketDST {
	return &cacherMarketDST{
		bean,
		mdCache,
	}
}

// implement Cacheor1
func (c *cacherMarketDST) Insert() {
	return
}

func (c *cacherMarketDST) Delete() {
}

func (c *cacherMarketDST) Update() {
	c.cache.set(c.bean)
}

func (c *cacherMarketDST) Get() (i interface{}, exist bool) {
	return c.cache.get(c.bean)
}

func (c *cacherMarketDST) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *cacherMarketDST) Cache(i interface{}) {
	md := i.([]structure.MarketDST)
	for i := range md {
		c.cache.set(&md[i])
	}
}

// cache
func (c *marketDSTCache) set(marketDST *structure.MarketDST) {
	c.Lock()
	defer c.Unlock()

	c.info[marketDST.MarketOwnerType] = marketDST
}

func (c *marketDSTCache) get(marketDST *structure.MarketDST) (res interface{}, exist bool) {
	c.RLock()
	defer c.RUnlock()

	res, exist = c.info[marketDST.MarketOwnerType]

	return
}

func (c *marketDSTCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
