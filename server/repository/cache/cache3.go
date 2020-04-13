package cache

import (
	"configer/server/repository/cache/cache3"
	"configer/server/structure"
	"fmt"
)

type CacherHoliday struct {
	*Cacher3
}

type Cacher3 struct {
	bean structure.Cacheor3
	cache cache3.Cache3
}

func NewCacherHoliday(bean *structure.Holiday) *CacherHoliday {
	return &CacherHoliday{
		&Cacher3{
		bean,
		holiCache,
		},
	}
}

func (c *Cacher3) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *Cacher3) Delete() (num int64, err error) {
	c.cache.Delete(c.bean)
	return
}

func (c *Cacher3) Update() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *Cacher3) Get() (i interface{}, exist bool) {
	c.cache.Get(c.bean)
	return
}

func (c *Cacher3) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *CacherHoliday) Cache(i interface{}) {
	ho := i.([]structure.Holiday)
	for i := range ho {
		c.cache.Insert(&ho[i])
	}
}