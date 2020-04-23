package mysql

import (
	"configer/server/structure"
)

type tablerConGroupSecurity struct {
	*tabler
}

func NewTablertablerConGroupSecurity(bean *structure.ConGroupSec) *tablerConGroupSecurity {
	return &tablerConGroupSecurity{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func (t *tablerConGroupSecurity) Export() (interface{}, error) {
	i := []structure.ConGroupSec{}
	err := t.Table(t.bean).Where(t.bean.ExportCondition()).Find(&i)

	return i, err
}
