package nameIDor

import (
	"configer/server/structure"
	"sync"
)

type NameIDer struct {
	ID2Name map[int]string
	name2ID map[string]int
	info    map[string]structure.NameIDor
	sync.RWMutex
}

func NewNameIDer() *NameIDer {
	return &NameIDer{
		ID2Name: make(map[int]string),
		name2ID: make(map[string]int),
		info:    make(map[string]structure.NameIDor),
	}
}

// cache
func (c *NameIDer) Insert(bean structure.NameIDor) {
	c.Lock()
	defer c.Unlock()

	ID := bean.GetID()
	name := bean.GetName()

	c.ID2Name[ID] = name
	c.name2ID[name] = ID
	c.info[name] = bean
}

func (c *NameIDer) Delete(bean structure.NameIDor) {
	c.Lock()
	defer c.Unlock()

	var ID int
	var name string

	ID = bean.GetID()
	if ID != 0 {
		name = c.ID2Name[ID]
	} else {
		name = bean.GetName()
		ID = c.name2ID[name]
	}

	delete(c.name2ID, name)
	delete(c.ID2Name, ID)
	delete(c.info, name)
}

func (c *NameIDer) Update(bean structure.NameIDor) {
	c.Lock()
	defer c.Unlock()

	var ID int
	var name string

	ID = bean.GetID()
	if ID != 0 {
		name = c.ID2Name[ID]
	} else {
		name = bean.GetName()
		ID = c.name2ID[name]
	}

	c.info[name] = bean
}

func (c *NameIDer) Get(bean structure.NameIDor) (res structure.NameIDor, exist bool) {
	c.RLock()
	defer c.RUnlock()

	var ID int
	var name string

	ID = bean.GetID()
	if ID != 0 {
		name = c.ID2Name[ID]
	} else {
		name = bean.GetName()
		ID = c.name2ID[name]
	}

	res, exist = c.info[name]

	return
}

func (c *NameIDer) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
