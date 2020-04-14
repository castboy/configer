package cache

import (
	"configer/server/structure"
	"fmt"
)

type cacherSession struct {
	bean  *structure.Session
	cache *sessionCache
}

func NewCacherSession(bean *structure.Session) *cacherSession {
	return &cacherSession{
		bean,
		sessCache[int(bean.Dst*2)+int(bean.Type)],
	}
}

// implement Cacheor1
func (c *cacherSession) Insert() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *cacherSession) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *cacherSession) Update() (num int64, err error) {
	c.cache.set(c.bean)
	return
}

func (c *cacherSession) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *cacherSession) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *cacherSession) Cache(i interface{}) {
	se := i.([]structure.Session)
	for i := range se {
		c.cache.set(&se[i])
	}
}

// cache
func (c *sessionCache) set(session *structure.Session) {
	c.Lock()
	defer c.Unlock()

	c.info[session.SourceID] = session.Session
}

func (c *sessionCache) get(session *structure.Session) {
}

func (c *sessionCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
