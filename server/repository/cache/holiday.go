package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherHoliday struct {
	bean *structure.Holiday
	cache *holidayCache
}

func NewCacherHoliday(bean *structure.Holiday) *CacherHoliday {
	return &CacherHoliday{
		bean,
		holiCache,
	}
}

// implement Cacheor1
func (c *CacherHoliday) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherHoliday) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *CacherHoliday) Update() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *CacherHoliday) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *CacherHoliday) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *CacherHoliday) Cache(i interface{}) {
	ho := i.([]structure.Holiday)
	for i := range ho {
		c.cache.insert(&ho[i])
	}
}

// cache
func (c *holidayCache) insert(ho *structure.Holiday) {
	c.Lock()
	defer c.Unlock()

	c.info[ho.ID] = ho
}

func (c *holidayCache) delete(ho *structure.Holiday) {
	c.Lock()
	defer c.Unlock()

	delete(c.info, ho.ID)
}

func (c *holidayCache) get(ho *structure.Holiday) {
	c.RLock()
	defer c.RUnlock()

	ho = c.info[ho.ID]
}

func (c *holidayCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}