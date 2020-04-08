package check

import "configer/server/structure"

type CheckerConvSymbol struct {
	bean *structure.ConvSymbol
}

var checkerConvSymbol *CheckerConvSymbol

func NewCheckerConvSymbol(bean *structure.ConvSymbol) *CheckerConvSymbol {
	if checkerConvSymbol == nil {
		checkerConvSymbol = &CheckerConvSymbol{
			bean,
		}
	}

	return checkerConvSymbol
}

func (c *CheckerConvSymbol) FormatCheck() error {
	return c.bean.FormatCheck()
}
