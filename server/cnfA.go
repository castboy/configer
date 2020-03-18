package server

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

// Aer as an implement of interface defined in this package.

type Aer struct {
	a       *structure.A
	cacherA *cache.CacherA
	tablerA *mysql.TablerA
}

func NewAer(a *structure.A) *Aer {
	return &Aer{
		a,
		cache.GetCacherA(),
		mysql.GetTablerA(),
	}
}

func (c *Aer) tableInsert() (int64, error) {
	return c.tablerA.Insert(c.a)
}

func (c *Aer) tableDelete() (int64, error) {
	return c.tablerA.Delete(c.a)
}

func (c *Aer) tableUpdate() (int64, error) {
	return c.tablerA.Update(c.a)
}

func (c *Aer) tableGet() (bool, error) {
	return c.tablerA.Get(c.a)
}

func (c *Aer) cacheInsert() error {
	return c.cacherA.Insert(c.a)
}

func (c *Aer) cacheDelete() error {
	return c.cacherA.Delete(c.a)
}

func (c *Aer) cacheUpdate() error {
	return c.cacherA.Update(c.a)
}

func (c *Aer) cacheGet() bool {
	return c.cacherA.Get(c.a)
}

