package base

import (
	"configer/server/repository/cache"
	"configer/server/repository/mysql"
	"configer/server/structure"
)

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