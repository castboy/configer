package cache

import (
	cache "configer/server/repository/cache/indexName"
	structure2 "configer/server/structure"
	"configer/server/utils"
	"fmt"
)

type cacherConvSymbol struct {
	*namer
}

type namer struct {
	bean  structure2.Nameor
	cache cache.Nameor
}

func NewCacherConvSymbol(bean *structure2.ConvSymbol) *cacherConvSymbol {
	return &cacherConvSymbol{
		&namer{
			bean:  bean,
			cache: csCache[bean.ConvType],
		},
	}
}

// implement NameIDor
func (c *namer) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *namer) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *namer) Update() (num int64, err error) {
	return
}

func (c *namer) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *namer) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *cacherConvSymbol) Cache(i interface{}) {
	src := i.([]structure2.Source)
	for i := range src {
		bean := c.bean.(*structure2.ConvSymbol)

		bean.ConvInfo = utils.BuildConvInfo(src[i].Source, src)
		bean.SourceName = src[i].Source

		c.cache.Insert(bean)
	}

}
