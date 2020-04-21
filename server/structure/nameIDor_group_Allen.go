package structure

import (
	"configer/server/constant"
	"fmt"
	"github.com/juju/errors"
	"github.com/shopspring/decimal"
)

//Group
type AccountGroup struct {
	ID              int                     `xorm:"id autoincr"`
	Name            string                  `xorm:"name notnull"`
	DepositCurrency string                  `xorm:"deposit_currency notnull"`
	MarginStopOut   decimal.Decimal         `xorm:"margin_stop_out notnull"`
	MarginMode      constant.MarginCalcMode `xorm:"margin_mode notnull"`
	MarginCall      decimal.Decimal         `xorm:"margin_call"`
	TradeType       constant.GroupTradeType `xorm:trade_type notnull`
	IsChargeSwap    bool                    `xorm:"is_charge_swap"`
	Comment         string                  `xorm:"comment"`
}

func (grp *AccountGroup) FormatCheck() error {
	return nil
}

func (grp *AccountGroup) IndexCheck() error {
	if grp.GetID() == 0 && grp.GetName() == "" {
		return errors.NotValidf("ID and name")
	}

	return nil
}

func (grp *AccountGroup) NotFoundError() error {
	return errors.NotFoundf("AccountGroup, SecurityName: %s, ID: %d", grp.Name, grp.ID)
}

func (grp *AccountGroup) AutoCondition() (cond string) {
	if grp.Name != "" {
		cond = fmt.Sprintf("`name` = '%s'", grp.Name)
		return
	}

	if grp.ID != 0 {
		cond = fmt.Sprintf("`id` = %d", grp.ID)
		return
	}

	return
}

func (grp *AccountGroup) GetName() string {
	return grp.Name
}

func (grp *AccountGroup) GetID() int {
	return grp.ID
}
