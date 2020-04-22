package mysql

import (
	"configer/server/structure"
)

type tablerGroup struct {
	*tabler
}

func NewTablerAccountGroup(bean *structure.AccountGroup) *tablerGroup {
	return &tablerGroup{
		&tabler{
			bean,
			xEngine,
		},
	}
}

func (t *tablerGroup) Export() (interface{}, error) {
	i := []structure.AccountGroup{}
	err := t.Table(t.bean).Where(t.bean.ExportCondition()).Find(&i)

	return i, err
}
