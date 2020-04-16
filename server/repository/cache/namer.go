package cache

import (
	"configer/server/repository/cache/nameor"
	"configer/server/structure"
	"configer/server/utils"
	"fmt"
)

type cacherConvSymbol struct {
	*namer
}

type namer struct {
	bean  structure.Nameor
	cache nameor.Nameor
}

func NewCacherConvSymbol(bean *structure.ConvSymbol) *cacherConvSymbol {
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

func (c *cacherConvSymbol) Cache(j interface{}) {
	src := j.([]structure.Source)
	for i := range src {
		insertBean := &structure.ConvSymbol{}
		bean := c.bean.(*structure.ConvSymbol)
		if bean.ConvType == structure.MarginConv {
			insertBean.ConvInfo = utils.BuildConvInfo(src[i].MarginCurrency, src)
		} else if bean.ConvType == structure.ProfitConv {
			insertBean.ConvInfo = utils.BuildConvInfo(src[i].ProfitCurrency, src)
		}

		insertBean.SourceName = src[i].Source
		insertBean.ConvType = bean.ConvType

		c.cache.Insert(insertBean)
	}

}
