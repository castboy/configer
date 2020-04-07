package cache

import "configer/server/structure"

type CacherSource struct {
	bean  *structure.Source
	cache *sourceCache
}

var cacherSource *CacherSource

func NewCacherSource(bean *structure.Source) *CacherSource {
	if cacherSource == nil {
		cacherSource = &CacherSource{
			bean,
			srcCache,
		}
	}

	return cacherSource
}

// implement Cacheor
func (c *CacherSource) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherSource) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *CacherSource) Update() (num int64, err error) {
	c.cache.update(c.bean)
	return
}

func (c *CacherSource) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherSource) Export() (i interface{}, err error) {
	return
}

func (c *CacherSource) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *sourceCache) insert(source *structure.Source) {
	c.Lock()
	defer c.Unlock()

	c.ID2Name[source.ID] = source.Source
	c.name2ID[source.Source] = source.ID
	c.info[source.Source] = source
}

func (c *sourceCache) delete(source *structure.Source) {
	c.Lock()
	defer c.Unlock()

	if source.Source == "" {
		source.Source = c.ID2Name[source.ID]
	} else if source.ID == 0 {
		source.ID = c.name2ID[source.Source]
	}

	delete(c.name2ID, source.Source)
	delete(c.ID2Name, source.ID)
	delete(c.info, source.Source)
}

func (c *sourceCache) update(source *structure.Source) {
	c.Lock()
	defer c.Unlock()

	if source.Source == "" {
		source.Source = c.ID2Name[source.ID]
	} else if source.ID == 0 {
		source.ID = c.name2ID[source.Source]
	}

	// update fields obviously.
	c.info[source.Source].SourceType = source.SourceType
	c.info[source.Source].Digits = source.Digits
	c.info[source.Source].Multiply = source.Multiply
	c.info[source.Source].ContractSize = source.ContractSize
	c.info[source.Source].StopsLevel = source.StopsLevel
	c.info[source.Source].MarginMode = source.MarginMode
	c.info[source.Source].ProfitCurrency = source.ProfitCurrency
	//without ProfitMode
	c.info[source.Source].Currency = source.Currency
	c.info[source.Source].MarginCurrency = source.MarginCurrency
	c.info[source.Source].SwapType = source.SwapType
	// without SwapCurrency
	c.info[source.Source].SwapLong = source.SwapLong
	c.info[source.Source].SwapShort = source.SwapShort
	c.info[source.Source].Swap3Day = source.Swap3Day
	c.info[source.Source].MarketOwnerType = source.MarketOwnerType
}

func (c *sourceCache) get(source *structure.Source) {
	c.RLock()
	defer c.RUnlock()

	if source.Source == "" {
		source.Source = c.ID2Name[source.ID]
	} else if source.ID == 0 {
		source.ID = c.name2ID[source.Source]
	}

	source = c.info[source.Source]
}