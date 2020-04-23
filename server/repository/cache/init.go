package cache

import (
	"configer/server/repository/cache/idor"
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

type holidayCache struct {
	info map[string]map[int]*structure.Holiday
	sync.RWMutex
}

//conGroupSecurity
type groupSecurityCache struct {
	info map[int]map[int]*structure.ConGroupSec // key: groupID, key2: securityID
	sync.RWMutex
}

const AllSessionTypeLength = int(structure.SessionTypeLength) * int(structure.DSTTypeLength)

var symbCache *nameIDor.NameIDer
var srcCache *nameIDor.NameIDer
var secCache *nameIDor.NameIDer

var sessCache [AllSessionTypeLength]*idor.IDer
var mdCache *marketDSTCache
var fsnCache *fullSymbolNameCache
var csCache [structure.ConvTypeLength]*idor.IDer
var holiCache *holidayCache
var grpCache *nameIDor.NameIDer
var grpSecCache *groupSecurityCache

func init() {
	symbCache = nameIDor.NewNameIDer()
	srcCache = nameIDor.NewNameIDer()
	secCache = nameIDor.NewNameIDer()
	grpCache = nameIDor.NewNameIDer()

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
		csCache[i] = idor.NewIDer()
	}

	holiCache = &holidayCache{
		info: map[string]map[int]*structure.Holiday{},
	}

	grpSecCache = &groupSecurityCache{
		info: map[int]map[int]*structure.ConGroupSec{},
	}
}
