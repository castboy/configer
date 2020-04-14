package cache

import (
	"configer/server/repository/cache/indexID"
	"configer/server/repository/cache/indexName"
	"configer/server/repository/cache/indexNameID"
	"configer/server/structure"
	cache12 "configer/server/structure/indexNameID"
	"sync"
)

type sessionCache struct {
	info map[int]map[int32][]string // int -> sourceID; string -> timeSpan.
	sync.RWMutex
}

type marketDSTCache struct {
	info map[cache12.MarketType]*structure.MarketDST
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

const AllTypeLength int = 6
const ConvTypeLength int = 2

var symbCache *indexNameID.BaseCache
var srcCache *indexNameID.BaseCache
var secCache *indexNameID.BaseCache

var sessCache [AllTypeLength]*sessionCache
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [ConvTypeLength]*indexName.BaseCache
var holiCache *indexID.BaseCache
var holiCalcCache *holidayCalcCache

func init() {
	symbCache = indexNameID.NewBaseCache()
	srcCache = indexNameID.NewBaseCache()
	secCache = indexNameID.NewBaseCache()

	for i := 0; i < AllTypeLength; i++ {
		sessCache[i] = &sessionCache{
			info: make(map[int]map[int32][]string),
		}
	}

	mdCache = &marketDSTCache{
		info: make(map[cache12.MarketType]*structure.MarketDST),
	}

	fsnCache = &fullSymbolNameCache{
		info: make(map[structure.SymbolLeverage]*structure.FullSymbolName),
	}

	for i := 0; i < ConvTypeLength; i++ {
		csCache[i] = indexName.NewBaseCache()
	}

	holiCache = indexID.NewBaseCache()

	holiCalcCache = &holidayCalcCache{
		info: make(map[int]map[structure.DateSymbol]*structure.TimeSpan),
	}
}
