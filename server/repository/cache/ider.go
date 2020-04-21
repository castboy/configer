package cache

import (
	"configer/server/repository/cache/idor"
	"configer/server/structure"
	"configer/server/utils"
)

type cacherSession struct {
	*ider
}

type cacherConvSymbol struct {
	*ider
}

type ider struct {
	bean  structure.IDor
	cache idor.IDor
}

func NewCacherSession(bean *structure.Session) *cacherSession {
	return &cacherSession{
		&ider{
			bean,
			sessCache[int(bean.Dst*2)+int(bean.Type)],
		},
	}
}

func NewCacherConvSymbol(bean *structure.ConvSymbol) *cacherConvSymbol {
	return &cacherConvSymbol{
		&ider{
			bean:  bean,
			cache: csCache[bean.ConvType],
		},
	}
}

func (c *ider) Insert() {
	c.cache.Insert(c.bean)
}

func (c *ider) Delete() {
	c.cache.Delete(c.bean)
}

func (c *ider) Update() {
}

func (c *ider) Get() (i interface{}, exist bool) {
	return c.cache.Get(c.bean)
}

func (c *ider) Export() (i interface{}, err error) {
	return c.cache.Export()
}

func (c *cacherSession) Update() {
	ses := c.bean.(*structure.Session)
	ses.Session = utils.TrimZero(ses.Session)

	c.cache.Update(ses)
}

func (c *cacherSession) Cache(i interface{}) {
	se := i.([]structure.Session)
	for i := range se {
		c.cache.Update(&se[i])
	}
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

		insertBean.SourceID = src[i].ID
		insertBean.ConvType = bean.ConvType

		c.cache.Insert(insertBean)
	}

}
