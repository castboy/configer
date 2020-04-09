package cache

import (
	"configer/server/structure"
	"fmt"
)

type CacherHoliday struct {
	bean *structure.Holiday
	cache *holidayCache
}


var cacherHoliday *CacherHoliday

func NewCacherHoliday(bean *structure.Holiday) *CacherHoliday {
	if cacherHoliday == nil {
		cacherHoliday = &CacherHoliday{
			bean,
			holiCache,
		}
	}

	return cacherHoliday
}

// implement Cacheor
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

func (c *CacherHoliday) Get() (exist bool, err error) {
	c.cache.get(c.bean)
	return
}

func (c *CacherHoliday) Export() (i interface{}, err error) {
	return
}

func (c *CacherHoliday) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *holidayCache) insert(ho *structure.Holiday) {
	c.Lock()
	defer c.Unlock()

	if c.info[ho.DateSymbol] == nil {
		c.info[ho.DateSymbol] = ho.HolidayTime
		return
	}

	c.info[ho.DateSymbol] = append(c.info[ho.DateSymbol], ho.HolidayTime...)
}

func (c *holidayCache) delete(ho *structure.Holiday) {
	c.Lock()
	defer c.Unlock()

	delete(c.info, ho.DateSymbol)
}

func (c *holidayCache) get(ho *structure.Holiday) {
	c.RLock()
	defer c.RUnlock()

	ho.HolidayTime = c.info[ho.DateSymbol]
}