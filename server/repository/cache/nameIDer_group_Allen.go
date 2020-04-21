package cache

import (
	"configer/server/structure"
)

type cacherGroup struct {
	*nameIDer
}

func NewCacherAccountGroup(bean *structure.AccountGroup) *cacherGroup {
	return &cacherGroup{
		&nameIDer{
			bean,
			grpCache,
		},
	}
}

func (c *cacherGroup) Cache(i interface{}) {
	gps := i.([]structure.AccountGroup)
	for i := range gps {
		c.cache.Insert(&gps[i])
	}
}
