package cache

import (
	"configer/server/structure"
	"sync"
)

//conGroupSecurity
type groupSecurityCache struct {
	info map[int]map[int]*structure.ConGroupSec // key: groupID, key2: securityID
	sync.RWMutex
}

var grpSecCache *groupSecurityCache

