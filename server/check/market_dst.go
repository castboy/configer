package check

import "configer/server/structure"

type CheckerMarketDST struct {
	bean *structure.MarketDST
}

var checkerMarketDST *CheckerMarketDST

func NewCheckerMarketDST(bean *structure.MarketDST) *CheckerMarketDST {
	if checkerMarketDST == nil {
		checkerMarketDST = &CheckerMarketDST{
			bean,
		}
	}

	return checkerMarketDST
}

func (c *CheckerMarketDST) FormatCheck() error {
	return c.bean.FormatCheck()
}

