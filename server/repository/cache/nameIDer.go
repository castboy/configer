package cache

import (
	"configer/server/repository/cache/nameIDor"
	"configer/server/structure"
)

type cacherSymbol struct {
	*nameIDer
}

type cacherSource struct {
	*nameIDer
}

type cacherSecurity struct {
	*nameIDer
}

type cacherGroup struct {
	*nameIDer
}

type nameIDer struct {
	bean  structure.NameIDor
	cache nameIDor.NameIDor
}

func NewCacherSymbol(bean *structure.Symbol) *cacherSymbol {
	return &cacherSymbol{
		&nameIDer{
			bean,
			symbCache,
		},
	}
}

func NewCacherSource(bean *structure.Source) *cacherSource {
	return &cacherSource{
		&nameIDer{
			bean,
			srcCache,
		},
	}
}

func NewCacherSecurity(bean *structure.Security) *cacherSecurity {
	return &cacherSecurity{
		&nameIDer{
			bean,
			secCache,
		},
	}
}

// implement NameIDor
func (c *nameIDer) Insert() {
	c.cache.Insert(c.bean)
	return
}

func (c *nameIDer) Delete() {
	c.cache.Delete(c.bean)
}

func (c *nameIDer) Update() {
	c.cache.Update(c.bean)
}

func (c *nameIDer) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *nameIDer) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *cacherSymbol) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		c.cache.Insert(&sb[i])
	}
}

func (c *cacherSource) Cache(i interface{}) {
	src := i.([]structure.Source)
	for i := range src {
		c.cache.Insert(&src[i])
	}
}

func (c *cacherSecurity) Cache(i interface{}) {
	se := i.([]structure.Security)
	for i := range se {
		c.cache.Insert(&se[i])
	}
}

func NewCacherAccountGroup(bean *structure.AccountGroup) *cacherGroup {
	return &cacherGroup{
		&nameIDer{
			bean,
			grpCache,
		},
	}
}

func (c *cacherGroup) Cache(i interface{}) {
	gps := i.([]structure.AccountGroup)
	for i := range gps {
		c.cache.Insert(&gps[i])
	}
}
