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


var symbCache *symbolCache
var srcCache *sourceCache

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
}
