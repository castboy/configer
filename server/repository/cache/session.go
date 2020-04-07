package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherSession struct {
	bean  *structure.Session
	cache *sessionCache
}

var cacherSession *CacherSession

func NewCacherSession(bean *structure.Session) *CacherSession {
	if cacherSession == nil {
		cacherSession = &CacherSession{
			bean,
			sessCache[int(bean.Dst*2)+int(bean.Type)],
		}
	}

	return cacherSession
}

// implement Cacheor
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

func (c *CacherSession) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherSession) Export() (i interface{}, err error) {
	return
}

func (c *CacherSession) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *sessionCache) set(session *structure.Session) {
}

func (c *sessionCache) get(session *structure.Session) {
}