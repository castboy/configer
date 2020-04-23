package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

type baser struct {
	cacher cache.BaseOperator
	tabler mysql.BaseOperator
	helper structure.Helpor
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

func NewGrouper(bean *structure.AccountGroup) *baser {
	return &baser{
		cache.NewCacherAccountGroup(bean),
		mysql.NewTablerAccountGroup(bean),
		bean,
	}
}

func NewConGroupSec(bean *structure.ConGroupSec) *baser {
	return &baser{
		cache.NewCacherConGroupSeucrity(bean),
		mysql.NewTablertablerConGroupSecurity(bean),
		bean,
	}
}

func (a *baser) GetCacher() cache.BaseOperator {
	return a.cacher
}

func (a *baser) GetTabler() mysql.BaseOperator {
	return a.tabler
}

func (a *baser) GetHelper() structure.Helpor {
	return a.helper
}
