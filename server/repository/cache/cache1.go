package cache

import (
	"configer/server/repository/cache/cache1"
	"configer/server/structure"
)

type CacherSymbol struct {
	*Cacher
}

type CacherSource struct {
	*Cacher
}

type CacherSecurity struct {
	*Cacher
}

type Cacher struct {
	bean  structure.Cacheor1
	cache cache1.Cache1
}

func NewCacherSymbol(bean *structure.Symbol) *CacherSymbol {
	return &CacherSymbol{
		&Cacher{
			bean,
			symbCache,
		},
	}
}

func NewCacherSource(bean *structure.Source) *CacherSource {
	return &CacherSource{
		&Cacher{
			bean,
			srcCache,
		},
	}
}

func NewCacherSecurity(bean *structure.Security) *CacherSecurity {
	return &CacherSecurity{
		&Cacher{
			bean,
			secCache,
		},
	}
}

// implement Cacheor1
func (c *Cacher) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *Cacher) Delete() (num int64, err error) {
	c.cache.Delete(c.bean)
	return
}

func (c *Cacher) Update() (num int64, err error) {
	c.cache.Update(c.bean)
	return
}

func (c *Cacher) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *Cacher) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *CacherSymbol) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		c.cache.Insert(&sb[i])
	}
}

func (c *CacherSource) Cache(i interface{}) {
	src := i.([]structure.Source)
	for i := range src {
		c.cache.Insert(&src[i])
	}
}

func (c *CacherSecurity) Cache(i interface{}) {
	se := i.([]structure.Security)
	for i := range se {
		c.cache.Insert(&se[i])
	}
}
