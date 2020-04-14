package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type baser struct {
	cacher  cache.BaseOperator
	tabler  mysql.BaseOperator
	checker structure.Checkor
}

func NewHolidayer(bean *structure.Holiday) *baser {
	return &baser{
		cache.NewCacherHoliday(bean),
		mysql.NewTablerHoliday(bean),
		bean,
	}
}

func NewMarketDSTer(bean *structure.MarketDST) *baser {
	return &baser{
		cache.NewCacherMarketDST(bean),
		mysql.NewTablerMarketDST(bean),
		bean,
	}
}

func NewSecurityer(bean *structure.Security) *baser {
	return &baser{
		cache.NewCacherSecurity(bean),
		mysql.NewTablerSecurity(bean),
		bean,
	}
}

func NewSessioner(bean *structure.Session) *baser {
	return &baser{
		cache.NewCacherSession(bean),
		mysql.NewTablerSession(bean),
		bean,
	}
}

func NewSourcer(bean *structure.Source) *baser {
	return &baser{
		cache.NewCacherSource(bean),
		mysql.NewTablerSource(bean),
		bean,
	}
}

func NewSymboler(a *structure.Symbol) *baser {
	return &baser{
		cache.NewCacherSymbol(a),
		mysql.NewTablerSymbol(a),
		a,
	}
}

func (a *baser) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *baser) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *baser) GetChecker() structure.Checkor {
	return a.checker
}
