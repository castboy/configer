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

var symbCache *symbolCache

func init() {
	symbCache = &symbolCache{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[int]*structure.Symbol),
	}
}
