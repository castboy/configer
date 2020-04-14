package cache

import (
	cache "configer/server/repository/cache/indexID"
	structure "configer/server/structure/indexID"
	"fmt"
)

type CacherHoliday struct {
	*IndexID
}

type IndexID struct {
	bean  structure.IDor
	cache cache.IDor
}

func NewCacherHoliday(bean *structure.Holiday) *CacherHoliday {
	return &CacherHoliday{
		&IndexID{
			bean,
			holiCache,
		},
	}
}

func (c *IndexID) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *IndexID) Delete() (num int64, err error) {
	c.cache.Delete(c.bean)
	return
}

func (c *IndexID) Update() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *IndexID) Get() (i interface{}, exist bool) {
	c.cache.Get(c.bean)
	return
}

func (c *IndexID) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *CacherHoliday) Cache(i interface{}) {
	ho := i.([]structure.Holiday)
	for i := range ho {
		c.cache.Insert(&ho[i])
	}
}
