package cache

import (
	cache "configer/server/repository/cache/indexName"
	structure "configer/server/structure/indexName"
	"configer/server/structure/indexNameID"
	"configer/server/utils"
	"fmt"
)

type CacherConvSymbol struct {
	*IndexName
}

type IndexName struct {
	bean  structure.Nameor
	cache cache.Nameor
}

func NewCacherConvSymbol(bean *structure.ConvSymbol) *CacherConvSymbol {
	return &CacherConvSymbol{
		&IndexName{
			bean:  bean,
			cache: csCache[bean.ConvType],
		},
	}
}

// implement NameIDor
func (c *IndexName) Insert() (num int64, err error) {
	c.cache.Insert(c.bean)
	return
}

func (c *IndexName) Delete() (num int64, err error) {
	err = fmt.Errorf("Method Not Support!")
	return
}

func (c *IndexName) Update() (num int64, err error) {
	return
}

func (c *IndexName) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *IndexName) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *CacherConvSymbol) Cache(i interface{}) {
	src := i.([]indexNameID.Source)
	for i := range src {
		bean := c.bean.(*structure.ConvSymbol)

		bean.ConvInfo = utils.BuildConvInfo(src[i].Source, src)
		bean.SourceName = src[i].Source

		c.cache.Insert(bean)
	}

}
