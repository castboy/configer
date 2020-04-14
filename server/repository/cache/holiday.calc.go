package cache

import (
	"configer/server/structure"
	"configer/server/structure/indexNameID"
)

type cacherHolidayCalc struct {
	bean  *structure.HolidayCalc
	cache *holidayCalcCache
}

func NewCacherHolidayCalc(bean *structure.HolidayCalc) *cacherHolidayCalc {
	return &cacherHolidayCalc{
		bean,
		holiCalcCache,
	}
}

// implement NameIDor
func (c *cacherHolidayCalc) Insert() (num int64, err error) {
	c.cache.insert(c.bean)
	return
}

func (c *cacherHolidayCalc) Delete() (num int64, err error) {
	c.cache.delete(c.bean)
	return
}

func (c *cacherHolidayCalc) Update() (num int64, err error) {
	c.cache.update(c.bean)
	return
}

func (c *cacherHolidayCalc) Get() (i interface{}, exist bool) {
	c.cache.get(c.bean)
	return
}

func (c *cacherHolidayCalc) Export() (i interface{}, err error) {
	return
}

func (c *cacherHolidayCalc) Cache(i interface{}) {
	sb := i.([]indexNameID.Symbol)
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
