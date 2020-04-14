package idor

import (
	"configer/server/structure"
	"sync"
)

type IDer struct {
	info map[int]structure.IDor
	sync.RWMutex
}

func NewIDer() *IDer {
	return &IDer{
		info: map[int]structure.IDor{},
	}
}

// cache
func (c *IDer) Insert(bean structure.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *IDer) Delete(bean structure.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	delete(c.info, ID)
}

func (c *IDer) Update(bean structure.IDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	c.info[ID] = bean
}

func (c *IDer) Get(bean structure.IDor) (res structure.IDor, exist bool) {
	c.RLock()
	defer c.RUnlock()

	ID := bean.GetID()

	res, exist = c.info[ID]
	return
}

func (c *IDer) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
