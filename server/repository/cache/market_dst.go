package cache

import (
	"configer/server/structure"
	"fmt"
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
func (c *cacherMarketDST) Insert() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *cacherMarketDST) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *cacherMarketDST) Update() (num int64, err error) {
	c.cache.set(c.bean)
	return
}

func (c *cacherMarketDST) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
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

func (c *marketDSTCache) get(marketDST *structure.MarketDST) {
	c.RLock()
	defer c.RUnlock()

	marketDST = c.info[marketDST.MarketOwnerType]
}

func (c *marketDSTCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
