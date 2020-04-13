package cache

import (
	"configer/server/structure"
	"sync"
)

type baseCache struct {
	ID2Name map[int]string
	name2ID map[string]int
	info    map[string]structure.Cacheor
	sync.RWMutex
}

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

type convSymbolCache struct {
	info map[string]*structure.ConvSymbol
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

var symbCache *baseCache
var srcCache *baseCache
var secCache *baseCache

var sessCache [AllTypeLength]*sessionCache
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [ConvTypeLength]*convSymbolCache
var holiCache *holidayCache
var holiCalcCache *holidayCalcCache

func init() {
	symbCache = &baseCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]structure.Cacheor),
	}

	srcCache = &baseCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]structure.Cacheor),
	}

	secCache = &baseCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]structure.Cacheor),
	}

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
		csCache[i] = &convSymbolCache{
			info: make(map[string]*structure.ConvSymbol),
		}
	}

	holiCache = &holidayCache{
		info: make(map[int]*structure.Holiday),
	}

	holiCalcCache = &holidayCalcCache{
		info: make(map[int]map[structure.DateSymbol]*structure.TimeSpan),
	}
}
