package cache

import (
	cache "configer/server/repository/cache/indexID"
	structure2 "configer/server/structure"
	"fmt"
)

type cacherHoliday struct {
	*ider
}

type cacherSession struct {
	*ider
}

type ider struct {
	bean  structure2.IDor
	cache cache.IDor
}

func NewCacherHoliday(bean *structure2.Holiday) *cacherHoliday {
	return &cacherHoliday{
		&ider{
			bean,
			holiCache,
		},
	}
}

func NewCacherSession(bean *structure2.Session) *cacherSession {
	return &cacherSession{
		&ider{
		bean,
		sessCache[int(bean.Dst*2)+int(bean.Type)],
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
	ho := i.([]structure2.Holiday)
	for i := range ho {
		c.cache.Insert(&ho[i])
	}
}

func (c *cacherSession) Cache(i interface{}) {
	se := i.([]structure2.Session)
	for i := range se {
		c.cache.Update(&se[i])
	}
}
