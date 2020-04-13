package cache

import (
	"configer/server/repository/cache/cache2"
	"configer/server/structure"
	"configer/server/utils"
	"fmt"
)

type CacherConvSymbol struct {
	*Cacher2
}

type Cacher2 struct {
	bean structure.Cacheor2
	cache cache2.Cache2
}

func NewCacherConvSymbol(bean *structure.ConvSymbol) *CacherConvSymbol {
	return &CacherConvSymbol{
		&Cacher2{
			bean: bean,
			cache: csCache[bean.ConvType],
		},
	}
}

// implement Cacheor1
func (c *Cacher2) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *Cacher2) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *Cacher2) Update() (num int64, err error) {
	return
}

func (c *Cacher2) Get() (exist bool, err error) {
	c.cache.Get(c.bean)
	return
}

func (c *Cacher2) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *CacherConvSymbol) Cache(i interface{}) {
	src := i.([]structure.Source)
	for i := range src {
		bean := c.bean.(*structure.ConvSymbol)

		bean.ConvInfo = utils.BuildConvInfo(src[i].Source, src)
		bean.SourceName = src[i].Source

		c.cache.Insert(bean)
	}

}