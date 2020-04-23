package cache

import (
	"configer/server/structure"
)

type cacherConGroupSeucrity struct {
	bean  *structure.ConGroupSec
	cache *groupSecurityCache
}

func NewCacherConGroupSeucrity(bean *structure.ConGroupSec) *cacherConGroupSeucrity {
	return &cacherConGroupSeucrity{
		bean,
		grpSecCache,
	}
}

func (c *cacherConGroupSeucrity) Insert() {
	c.cache.insert(c.bean)
}

func (c *cacherConGroupSeucrity) Delete() {
	c.cache.delete(c.bean.GroupId, c.bean.SecurityId)
}

func (c *cacherConGroupSeucrity) Update() {
	c.cache.update(c.bean)
}

func (c *cacherConGroupSeucrity) Get() (i interface{}, exist bool) {
	return c.cache.get(c.bean.GroupId, c.bean.SecurityId)
}

func (c *cacherConGroupSeucrity) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *cacherConGroupSeucrity) Cache(i interface{}) {
	cgss := i.([]structure.ConGroupSec)
	for i := range cgss {
		c.cache.insert(&cgss[i])
	}
}

// cache

func (c *groupSecurityCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}

//ByGroupIDAndSecurityID
func (c *groupSecurityCache) get(groupID, securityID int) (res *structure.ConGroupSec, exist bool) {
	c.RLock()
	defer c.RUnlock()

	res, exist = c.info[groupID][securityID]

	return
}

func (c *groupSecurityCache) insert(bean *structure.ConGroupSec) {
	c.Lock()
	defer c.Unlock()

	if c.info[bean.GroupId] == nil {
		c.info[bean.GroupId] = map[int]*structure.ConGroupSec{}
	}

	c.info[bean.GroupId][bean.SecurityId] = bean
}

//updateByGroupIDAndSecurityID
func (c *groupSecurityCache) update(gs *structure.ConGroupSec) {
	c.insert(gs)
}

//deleteByGroupIDAndSecurityID
func (c *groupSecurityCache) delete(groupID, securityID int) error {
	c.Lock()
	defer c.Unlock()

	delete(c.info[groupID], securityID)

	if len(c.info[groupID]) == 0 {
		delete(c.info, groupID)
	}

	return nil
}
