package cache

import (
	"configer/server/structure"
)

type cacherHoliday struct {
	bean  *structure.Holiday
	cache *holidayCache
}

func NewCacherHoliday(bean *structure.Holiday) *cacherHoliday {
	return &cacherHoliday{
		bean,
		holiCache,
	}
}

func (c *cacherHoliday) Insert() {
	c.cache.insert(c.bean)
}

func (c *cacherHoliday) Delete() {
	c.cache.delete(c.bean)
}

func (c *cacherHoliday) Update() {
	c.cache.update(c.bean)
}

func (c *cacherHoliday) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *cacherHoliday) Export() (i interface{}, err error) {
	return c.cache.export()
}

func (c *cacherHoliday) Cache(i interface{}) {
	sb := i.([]structure.Holiday)
	for i := range sb {
		c.cache.insert(&sb[i])
	}
}

// cache
func (c *holidayCache) insert(bean *structure.Holiday) {
	c.Lock()
	defer c.Unlock()

	if c.info[bean.Date] == nil {
		c.info[bean.Date] = map[int]*structure.Holiday{}
	}

	c.info[bean.Date][bean.ID] = bean
}

func (c *holidayCache) delete(bean *structure.Holiday) {
	c.Lock()
	defer c.Unlock()

	deleted := false

	for date := range c.info {
		for id := range c.info[date] {
			if id == bean.ID {
				delete(c.info[date], bean.ID)
				deleted = true
				break
			}
		}

		if deleted {
			if len(c.info[date]) == 0 {
				delete(c.info, date)
			}

			return
		}
	}
}

func (c *holidayCache) get(bean *structure.Holiday) (i interface{}, exist bool) {
	c.RLock()
	defer c.RUnlock()

	i, exist = c.info[bean.Date]
	return
}

func (c *holidayCache) update(bean *structure.Holiday) {
	c.RLock()
	defer c.RUnlock()

	// if `Date` is not updated.
	if c.info[bean.Date][bean.ID] != nil {
		c.info[bean.Date][bean.ID] = bean
		return
	}

	// if `Date` is updated.
	c.RUnlock()

	c.delete(bean)
	c.insert(bean)
}

func (c *holidayCache) export() (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	return c.info, nil
}
