package cache

import (
	cache "configer/server/repository/cache/indexNameID"
	structure "configer/server/structure/indexNameID"
)

type CacherSymbol struct {
	*IndexNameID
}

type CacherSource struct {
	*IndexNameID
}

type CacherSecurity struct {
	*IndexNameID
}

type IndexNameID struct {
	bean  structure.NameIDor
	cache cache.NameIDor
}

func NewCacherSymbol(bean *structure.Symbol) *CacherSymbol {
	return &CacherSymbol{
		&IndexNameID{
			bean,
			symbCache,
		},
	}
}

func NewCacherSource(bean *structure.Source) *CacherSource {
	return &CacherSource{
		&IndexNameID{
			bean,
			srcCache,
		},
	}
}

func NewCacherSecurity(bean *structure.Security) *CacherSecurity {
	return &CacherSecurity{
		&IndexNameID{
			bean,
			secCache,
		},
	}
}

// implement NameIDor
func (c *IndexNameID) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *IndexNameID) Delete() (num int64, err error) {
	c.cache.Delete(c.bean)
	return
}

func (c *IndexNameID) Update() (num int64, err error) {
	c.cache.Update(c.bean)
	return
}

func (c *IndexNameID) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *IndexNameID) Export() (i interface{}, err error) {
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
