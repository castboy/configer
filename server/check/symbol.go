package check

import "configer/server/structure"

type CheckerSymbol struct {
	bean *structure.Symbol
}

var checkerSymbol *CheckerSymbol

func NewCheckerSymbol(bean *structure.Symbol) *CheckerSymbol {
	if checkerSymbol == nil {
		checkerSymbol = &CheckerSymbol{
			bean,
		}
	}

	return checkerSymbol
}

func (c *CheckerSymbol) FormatCheck() error {
	return c.bean.FormatCheck()
}