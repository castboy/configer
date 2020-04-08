package check

import "configer/server/structure"

type CheckerFullSymbolName struct {
	bean *structure.FullSymbolName
}

var checkerFullSymbolName *CheckerFullSymbolName

func NewCheckerFullSymbolName(bean *structure.FullSymbolName) *CheckerFullSymbolName {
	if checkerFullSymbolName == nil {
		checkerFullSymbolName = &CheckerFullSymbolName{
			bean,
		}
	}

	return checkerFullSymbolName
}

func (c *CheckerFullSymbolName) FormatCheck() error {
	return c.bean.FormatCheck()
}
