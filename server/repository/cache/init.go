package cache

import (
	"configer/server/repository/cache/cache1"
	"configer/server/repository/cache/cache2"
	"configer/server/structure"
	"sync"
)

type sessionCache struct {
	info map[int]map[int32][]string // int -> sourceID; string -> timeSpan.
	sync.RWMutex
}

type marketDSTCache struct {
	info map[structure.MarketType]*structure.MarketDST
	sync.RWMutex
}

type fullSymbolNameCache struct {
	info map[structure.SymbolLeverage]*structure.FullSymbolName
	sync.RWMutex
}

type holidayCache struct {
	info map[int]*structure.Holiday
	sync.RWMutex
}

type holidayCalcCache struct {
	info map[int]map[structure.DateSymbol]*structure.TimeSpan
	sync.RWMutex
}

const AllTypeLength int = 6
const ConvTypeLength int = 2

var symbCache *cache1.BaseCache
var srcCache *cache1.BaseCache
var secCache *cache1.BaseCache

var sessCache [AllTypeLength]*sessionCache
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [ConvTypeLength]*cache2.BaseCache
var holiCache *holidayCache
var holiCalcCache *holidayCalcCache

func init() {
	symbCache = cache1.NewBaseCache()
	srcCache = cache1.NewBaseCache()
	secCache = cache1.NewBaseCache()

	for i := 0; i < AllTypeLength; i++ {
		sessCache[i] = &sessionCache{
			info: make(map[int]map[int32][]string),
		}
	}

	mdCache = &marketDSTCache{
		info: make(map[structure.MarketType]*structure.MarketDST),
	}

	fsnCache = &fullSymbolNameCache{
		info: make(map[structure.SymbolLeverage]*structure.FullSymbolName),
	}

	for i := 0; i < ConvTypeLength; i++ {
		csCache[i] = cache2.NewBaseCache()
	}

	holiCache = &holidayCache{
		info: make(map[int]*structure.Holiday),
	}

	holiCalcCache = &holidayCalcCache{
		info: make(map[int]map[structure.DateSymbol]*structure.TimeSpan),
	}
}
