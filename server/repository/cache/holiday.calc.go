package cache

import (
	"configer/server/structure"
)

type CacherHolidayCalc struct {
	bean *structure.HolidayCalc
	cache *holidayCalcCache
}

func NewCacherHolidayCalc(bean *structure.HolidayCalc) *CacherHolidayCalc {
	return &CacherHolidayCalc{
		bean,
		holiCalcCache,
	}
}

// implement Cacheor1
func (c *CacherHolidayCalc) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *CacherHolidayCalc) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *CacherHolidayCalc) Update() (num int64, err error) {
	c.cache.update(c.bean)
	return
}

func (c *CacherHolidayCalc) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *CacherHolidayCalc) Export() (i interface{}, err error) {
	return
}

func (c *CacherHolidayCalc) Cache(i interface{}) {
	sb := i.([]structure.Symbol)
	for i := range sb {
		_ = i
	}
}

// cache
func (c *holidayCalcCache) insert(ho *structure.HolidayCalc) {
	c.Lock()
	defer c.Unlock()

	if c.info[ho.ID] == nil {
		c.info[ho.ID] = make(map[structure.DateSymbol]*structure.TimeSpan)
	}

	c.info[ho.ID][ho.DateSymbol] = ho.TimeSpans[0]
}

func (c *holidayCalcCache) delete(ho *structure.HolidayCalc) {
	c.Lock()
	defer c.Unlock()

	delete(c.info, ho.ID)
}

func (c *holidayCalcCache) update(ho *structure.HolidayCalc) {
	c.Lock()
	defer c.Unlock()

	for i := range c.info[ho.ID] {
		c.info[ho.ID][i] = ho.TimeSpans[0]
	}
}


func (c *holidayCalcCache) get(ho *structure.HolidayCalc) {
	c.RLock()
	defer c.RUnlock()

	ho.TimeSpans = []*structure.TimeSpan{}
	for i := range c.info {
		ts := c.info[i][ho.DateSymbol]
		if ts != nil {
			ho.TimeSpans = append(ho.TimeSpans, ts)
		}
	}
}