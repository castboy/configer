package check

import "configer/server/structure"

type CheckerHoliday struct {
	bean *structure.Holiday
}

var checkerHoliday *CheckerHoliday

func NewCheckerHoliday(bean *structure.Holiday) *CheckerHoliday {
	if checkerHoliday == nil {
		checkerHoliday = &CheckerHoliday{
			bean,
		}
	}

	return checkerHoliday
}

func (c *CheckerHoliday) FormatCheck() error {
	return c.bean.FormatCheck()
}
