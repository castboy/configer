package server

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

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

func (c *Aer) TableInsert() (int64, error) {
	return c.tablerA.Insert(c.a)
}

func (c *Aer) TableDelete() (int64, error) {
	return c.tablerA.Delete(c.a)
}

func (c *Aer) TableUpdate() (int64, error) {
	return c.tablerA.Update(c.a)
}

func (c *Aer) TableGet() (bool, error) {
	return c.tablerA.Get(c.a)
}

func (c *Aer) CacheInsert() error {
	return c.cacherA.Insert(c.a)
}

func (c *Aer) CacheDelete() error {
	return c.cacherA.Delete(c.a)
}

func (c *Aer) CacheUpdate() error {
	return c.cacherA.Update(c.a)
}

func (c *Aer) CacheGet() bool {
	return c.cacherA.Get(c.a)
}

