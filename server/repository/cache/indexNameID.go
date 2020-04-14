package cache

import (
	cache "configer/server/repository/cache/indexNameID"
	structure2 "configer/server/structure"
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

type nameIDer struct {
	bean  structure2.NameIDor
	cache cache.NameIDor
}

func NewCacherSymbol(bean *structure2.Symbol) *cacherSymbol {
	return &cacherSymbol{
		&nameIDer{
			bean,
			symbCache,
		},
	}
}

func NewCacherSource(bean *structure2.Source) *cacherSource {
	return &cacherSource{
		&nameIDer{
			bean,
			srcCache,
		},
	}
}

func NewCacherSecurity(bean *structure2.Security) *cacherSecurity {
	return &cacherSecurity{
		&nameIDer{
			bean,
			secCache,
		},
	}
}

// implement NameIDor
func (c *nameIDer) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *nameIDer) Delete() (num int64, err error) {
	c.cache.Delete(c.bean)
	return
}

func (c *nameIDer) Update() (num int64, err error) {
	c.cache.Update(c.bean)
	return
}

func (c *nameIDer) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *nameIDer) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *cacherSymbol) Cache(i interface{}) {
	sb := i.([]structure2.Symbol)
	for i := range sb {
		c.cache.Insert(&sb[i])
	}
}

func (c *cacherSource) Cache(i interface{}) {
	src := i.([]structure2.Source)
	for i := range src {
		c.cache.Insert(&src[i])
	}
}

func (c *cacherSecurity) Cache(i interface{}) {
	se := i.([]structure2.Security)
	for i := range se {
		c.cache.Insert(&se[i])
	}
}
