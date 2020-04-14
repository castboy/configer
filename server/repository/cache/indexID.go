package cache

import (
	cache "configer/server/repository/cache/indexID"
	structure "configer/server/structure/indexID"
	"fmt"
)

type cacherHoliday struct {
	*ider
}

type ider struct {
	bean  structure.IDor
	cache cache.IDor
}

func NewCacherHoliday(bean *structure.Holiday) *cacherHoliday {
	return &cacherHoliday{
		&ider{
			bean,
			holiCache,
		},
	}
}

func (c *ider) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *ider) Delete() (num int64, err error) {
	c.cache.Delete(c.bean)
	return
}

func (c *ider) Update() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *ider) Get() (i interface{}, exist bool) {
	c.cache.Get(c.bean)
	return
}

func (c *ider) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *cacherHoliday) Cache(i interface{}) {
	ho := i.([]structure.Holiday)
	for i := range ho {
		c.cache.Insert(&ho[i])
	}
}
