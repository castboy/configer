package nameor

import (
	"configer/server/structure"
	"sync"
)

type Namer struct {
	info map[string]structure.Nameor
	sync.RWMutex
}

func NewNamer() *Namer {
	return &Namer{
		info: map[string]structure.Nameor{},
	}
}

// cache
func (c *Namer) Insert(bean structure.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *Namer) Delete(bean structure.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	delete(c.info, name)
}

func (c *Namer) Update(bean structure.Nameor) {
	c.Lock()
	defer c.Unlock()

	name := bean.GetName()
	c.info[name] = bean
}

func (c *Namer) Get(bean structure.Nameor) (res structure.Nameor, exist bool) {
	c.RLock()
	defer c.RUnlock()

	name := bean.GetName()

	res, exist = c.info[name]
	return
}

func (c *Namer) Export() (i interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
