package cache

import (
	"configer/server/repository/cache/idor"
	"configer/server/repository/cache/nameor"
	"configer/server/repository/cache/nameIDor"
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

var symbCache *nameIDor.NameIDer
var srcCache *nameIDor.NameIDer
var secCache *nameIDor.NameIDer

var sessCache [AllSessionTypeLength]*idor.IDer
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [structure.ConvTypeLength]*nameor.Namer
var holiCache *idor.IDer
var holiCalcCache *holidayCalcCache

func init() {
	symbCache = nameIDor.NewNameIDer()
	srcCache = nameIDor.NewNameIDer()
	secCache = nameIDor.NewNameIDer()

	for i := 0; i < AllSessionTypeLength; i++ {
		sessCache[i] = idor.NewIDer()
	}

	mdCache = &marketDSTCache{
		info: make(map[structure.MarketType]*structure.MarketDST),
	}

	fsnCache = &fullSymbolNameCache{
		info: make(map[structure.SymbolLeverage]*structure.FullSymbolName),
	}

	for i := 0; i < int(structure.ConvTypeLength); i++ {
		csCache[i] = nameor.NewNamer()
	}

	holiCache = idor.NewIDer()

	holiCalcCache = &holidayCalcCache{
		info: make(map[int]map[structure.DateSymbol]*structure.TimeSpan),
	}
}
