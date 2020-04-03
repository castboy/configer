package check

import "configer/server/structure"

type CheckerSymbol struct {
	a *structure.Symbol
}

var checkerSymbol *CheckerSymbol

func NewCheckerSymbol(a *structure.Symbol) *CheckerSymbol {
	if checkerSymbol == nil {
		checkerSymbol = &CheckerSymbol{
			a,
		}
	}

	return checkerSymbol
}

func (c *CheckerSymbol) FormatCheck() error {
	return nil
}