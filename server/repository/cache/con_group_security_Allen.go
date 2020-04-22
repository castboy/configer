package cache

import (
	"configer/server/structure"
	"github.com/juju/errors"
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
	c.cache.delete(c.bean.GroupId,c.bean.SecurityId)
}

func (c *cacherConGroupSeucrity) Update() {
	c.cache.update(c.bean)
}

func (c *cacherConGroupSeucrity) Get() (i interface{}, exist bool) {
	res,err:=c.cache.get(c.bean.GroupId,c.bean.SecurityId)
	if err!=nil{
		return nil,false
	}
	return res,true
}

func (c *cacherConGroupSeucrity) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *cacherConGroupSeucrity) Cache(i interface{}) {
	cgss:= i.([]structure.ConGroupSec)
	for i := range cgss {
			c.cache.insert(&cgss[i])
	}
}


// cache

func (c *groupSecurityCache)export() (interface{},error){
	c.RLock()
	defer c.RUnlock()

	return c.info,nil
}
//ByGroupIDAndSecurityID
func (c *groupSecurityCache) get(groupID, securityID int) (*structure.ConGroupSec, error) {
	c.RLock()
	defer c.RUnlock()

	if c.info[groupID] == nil {
		return nil, errors.NotFoundf("group id: ", groupID)
	}

	if c.info[groupID][securityID] == nil {
		return nil, errors.NotFoundf("security id: ", securityID)
	}
	res := *(c.info[groupID][securityID])

	return &res, nil
}

func (c *groupSecurityCache) insert(bean *structure.ConGroupSec)  {

	c.Lock()
	defer c.Unlock()

	if c.info[bean.GroupId] == nil {
		c.info[bean.GroupId] = map[int]*structure.ConGroupSec{}
	}

	c.info[bean.GroupId][bean.SecurityId] = bean
}
//updateByGroupIDAndSecurityID
func (c *groupSecurityCache) update(gs *structure.ConGroupSec) error {
	if gs == nil {
		return errors.New("group-security is nil")
	}

	c.Lock()
	defer c.Unlock()

	if c.info[gs.GroupId] == nil {
		return errors.NotFoundf("group id: ", gs.GroupId)
	}

	if c.info[gs.GroupId][gs.SecurityId] == nil {
		return errors.NotFoundf("security id: ", gs.SecurityId)
	}

	// update fields obviously.
	c.info[gs.GroupId][gs.SecurityId].EnableSecurity = gs.EnableSecurity
	c.info[gs.GroupId][gs.SecurityId].EnableTrade = gs.EnableTrade
	c.info[gs.GroupId][gs.SecurityId].LotMin = gs.LotMin
	c.info[gs.GroupId][gs.SecurityId].LotMax = gs.LotMax
	c.info[gs.GroupId][gs.SecurityId].LotStep = gs.LotStep
	c.info[gs.GroupId][gs.SecurityId].SpreadDiff = gs.SpreadDiff
	c.info[gs.GroupId][gs.SecurityId].Commission = gs.Commission

	return nil
}
//deleteByGroupIDAndSecurityID
func (c *groupSecurityCache) delete(groupID, securityID int) error {
	c.Lock()
	defer c.Unlock()

	if c.info[groupID] == nil {
		return errors.NotFoundf("group id: %d", groupID)
	}

	if c.info[groupID][securityID] == nil {
		return errors.NotFoundf("security id: %d", securityID)
	}

	delete(c.info[groupID], securityID)

	if len(c.info[groupID]) == 0 {
		delete(c.info, groupID)
	}

	return nil
}

func (c *groupSecurityCache) isGroupHoldSecurity(groupID int) bool {
	c.RLock()
	defer c.RUnlock()

	if c.info[groupID] != nil && len(c.info[groupID]) != 0 {
		return true
	}

	return false
}