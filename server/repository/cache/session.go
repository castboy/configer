package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherSession struct {
	bean  *structure.Session
	cache *sessionCache
}

func NewCacherSession(bean *structure.Session) *CacherSession {
	return &CacherSession{
		bean,
		sessCache[int(bean.Dst*2)+int(bean.Type)],
	}
}

// implement Cacheor1
func (c *CacherSession) Insert() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherSession) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherSession) Update() (num int64, err error) {
	c.cache.set(c.bean)
	return
}

func (c *CacherSession) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *CacherSession) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *CacherSession) Cache(i interface{}) {
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