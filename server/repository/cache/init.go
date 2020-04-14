package cache

import (
	"configer/server/repository/cache/indexID"
	cache "configer/server/repository/cache/indexName"
	cache1 "configer/server/repository/cache/indexNameID"
	"configer/server/structure"
	"sync"
)

type marketDSTCache struct {
	info map[structure.MarketType]*structure.MarketDST
	sync.RWMutex
}

type fullSymbolNameCache struct {
	info map[structure.SymbolLeverage]*structure.FullSymbolName
	sync.RWMutex
}

type holidayCalcCache struct {
	info map[int]map[structure.DateSymbol]*structure.TimeSpan
	sync.RWMutex
}

const AllSessionTypeLength = int(structure.SessionTypeLength)*int(structure.DSTTypeLength)

var symbCache *cache1.BaseCache
var srcCache *cache1.BaseCache
var secCache *cache1.BaseCache

var sessCache [AllSessionTypeLength]*indexID.BaseCache
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [structure.ConvTypeLength]*cache.BaseCache
var holiCache *indexID.BaseCache
var holiCalcCache *holidayCalcCache

func init() {
	symbCache = cache1.NewBaseCache()
	srcCache = cache1.NewBaseCache()
	secCache = cache1.NewBaseCache()

	for i := 0; i < AllSessionTypeLength; i++ {
		sessCache[i] = indexID.NewBaseCache()
	}

	mdCache = &marketDSTCache{
		info: make(map[structure.MarketType]*structure.MarketDST),
	}

	fsnCache = &fullSymbolNameCache{
		info: make(map[structure.SymbolLeverage]*structure.FullSymbolName),
	}

	for i := 0; i < int(structure.ConvTypeLength); i++ {
		csCache[i] = cache.NewBaseCache()
	}

	holiCache = indexID.NewBaseCache()

	holiCalcCache = &holidayCalcCache{
		info: make(map[int]map[structure.DateSymbol]*structure.TimeSpan),
	}
}
