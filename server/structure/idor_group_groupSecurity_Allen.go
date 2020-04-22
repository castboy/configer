package structure

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/juju/errors"
)

//congroup security
//ConGroupSec
type ConGroupSec struct {
	ID             int             `xorm:"'id' pk " `
	GroupId        int             `xorm:"group_id"`
	SecurityId     int             `xorm:"security_id"`
	EnableSecurity bool            `xorm:"enable_security"`
	EnableTrade    bool            `xorm:"enable_trade"`
	LotMin         decimal.Decimal `xorm:"lot_min"`
	LotMax         decimal.Decimal `xorm:"lot_max"`
	LotStep        decimal.Decimal `xorm:"lot_step"`
	SpreadDiff     int             `xorm:"spread_diff"`
	Commission     decimal.Decimal `xorm:"commission"`
}

func (cgs *ConGroupSec) FormatCheck() error {
	if cgs == nil {
		return errors.NotValidf("session info is null")
	}

	if cgs.GroupId <=0  {
		return errors.NotValidf("groupID, %v", cgs.GroupId)
	}

	if cgs.SecurityId <=0 {
		return errors.NotValidf("SecurityId, %v", cgs.SecurityId)
	}

	if cgs.LotMin.LessThan(decimal.Zero)   {
		return errors.NotValidf("LotMin, %v", cgs.LotMin)
	}

	if cgs.LotMax.LessThan(decimal.Zero)   {
		return errors.NotValidf("LotMax, %v", cgs.LotMax)
	}

	if cgs.LotStep.LessThan(decimal.Zero)   {
		return errors.NotValidf("LotStep, %v", cgs.LotStep)
	}

	return nil
}

func (cgs *ConGroupSec) IndexCheck() error {
	if cgs.GetID() < 0 {
		return errors.NotValidf("ID")
	}

	return nil
}

func (cgs *ConGroupSec) NotFoundError() error {
	return errors.NotFoundf("ConGroupSec, groupID: %d, securityID:%d", cgs.GroupId,cgs.SecurityId)
}

func (cgs *ConGroupSec) ExportCondition() (cond string) {
	return "1"
}

func (cgs *ConGroupSec) UpdateCondition() (cond string) {
	if cgs.ID != 0 {
		return fmt.Sprintf("`id` = %d", cgs.ID)
	}

	return fmt.Sprintf("`group_id` = %d and `security_id` = %d", cgs.GroupId, cgs.SecurityId)
}

func (cgs *ConGroupSec) DeleteCondition() (cond string) {
	return cgs.UpdateCondition()
}

func (cgs *ConGroupSec) GetCondition() (cond string) {
	return cgs.UpdateCondition()
}

func (cgs *ConGroupSec) GetID() int {
	return cgs.ID
}
