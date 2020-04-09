package cache

import (
	"configer/server/structure"
	"sync"
)

type symbolCache struct {
	ID2Name map[int]string
	name2ID map[string]int
	info    map[int]*structure.Symbol
	sync.RWMutex
}

type sourceCache struct {
	ID2Name map[int]string
	name2ID map[string]int
	info    map[string]*structure.Source
	sync.RWMutex
}

type sessionCache struct {
	info map[int]map[int32][]string // int -> sourceID; string -> timeSpan.
	sync.RWMutex
}

type securityCache struct {
	ID2Name map[int]string
	name2ID map[string]int
	info    map[string]*structure.Security
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
	info map[structure.DateSymbol][]structure.HolidayTime
	sync.RWMutex
}

const AllTypeLength int = 6
const ConvTypeLength int = 2

var symbCache *symbolCache
var srcCache *sourceCache
var sessCache [AllTypeLength]*sessionCache
var secCache *securityCache
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [ConvTypeLength]*convSymbolCache
var holiCache *holidayCache

func init() {
	symbCache = &symbolCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[int]*structure.Symbol),
	}

	srcCache = &sourceCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]*structure.Source),
	}

	for i := 0; i < AllTypeLength; i++ {
		sessCache[i] = &sessionCache{
			info: make(map[int]map[int32][]string),
		}
	}

	secCache = &securityCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]*structure.Security),
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
		info: make(map[structure.DateSymbol][]structure.HolidayTime),
	}
}
