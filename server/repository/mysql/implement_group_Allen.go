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

func (t *tablerGroup) Update() (num int64, err error) {
	return t.Table(t.bean).Where(t.bean.AutoCondition()).Cols("id,name,deposit_currency,margin_stop_out,margin_mode,margin_call,trade_type,is_charge_swap,comment").Update(t.bean)
}

func (t *tablerGroup) Export() (interface{}, error) {
	i := []structure.AccountGroup{}
	err := t.Table(t.bean).Where(t.bean.AutoCondition()).Find(&i)

	return i, err
}
