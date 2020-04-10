package check

import "configer/server/structure"

type CheckerHolidayCalc struct {
	bean *structure.HolidayCalc
}

var checkerHolidayCalc *CheckerHolidayCalc

func NewCheckerHolidayCalc(bean *structure.HolidayCalc) *CheckerHolidayCalc {
	if checkerHolidayCalc == nil {
		checkerHolidayCalc = &CheckerHolidayCalc{
			bean,
		}
	}

	return checkerHolidayCalc
}

func (c *CheckerHolidayCalc) FormatCheck() error {
	return c.bean.FormatCheck()
}
