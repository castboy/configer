package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
	"configer/server/structure/indexID"
	"configer/server/structure/indexNameID"
)

type Baser struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker structure.Checkor
}

func NewHolidayer(bean *indexID.Holiday) *Baser {
	return &Baser{
		cache.NewCacherHoliday(bean),
		mysql.NewTablerHoliday(bean),
		bean,
	}
}

func NewMarketDSTer(bean *structure.MarketDST) *Baser {
	return &Baser{
		cache.NewCacherMarketDST(bean),
		mysql.NewTablerMarketDST(bean),
		bean,
	}
}

func NewSecurityer(bean *indexNameID.Security) *Baser {
	return &Baser{
		cache.NewCacherSecurity(bean),
		mysql.NewTablerSecurity(bean),
		bean,
	}
}

func NewSessioner(bean *structure.Session) *Baser {
	return &Baser{
		cache.NewCacherSession(bean),
		mysql.NewTablerSession(bean),
		bean,
	}
}

func NewSourcer(bean *indexNameID.Source) *Baser {
	return &Baser{
		cache.NewCacherSource(bean),
		mysql.NewTablerSource(bean),
		bean,
	}
}

func NewSymboler(a *indexNameID.Symbol) *Baser {
	return &Baser{
		cache.NewCacherSymbol(a),
		mysql.NewTablerSymbol(a),
		a,
	}
}

func (a *Baser) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *Baser) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *Baser) GetChecker() structure.Checkor {
	return a.checker
}
