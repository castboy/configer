package check

import "configer/server/structure"

type CheckerSecurity struct {
	bean *structure.Security
}

var checkerSecurity *CheckerSecurity

func NewCheckerSecurity(bean *structure.Security) *CheckerSecurity {
	if checkerSecurity == nil {
		checkerSecurity = &CheckerSecurity{
			bean,
		}
	}

	return checkerSecurity
}

func (c *CheckerSecurity) FormatCheck() error {
	return c.bean.FormatCheck()
}