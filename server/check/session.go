package check

import "configer/server/structure"

type CheckerSession struct {
	bean *structure.Session
}

var checkerSession *CheckerSession

func NewCheckerSession(bean *structure.Session) *CheckerSession {
	if checkerSession == nil {
		checkerSession = &CheckerSession{
			bean,
		}
	}

	return checkerSession
}

func (c *CheckerSession) FormatCheck() error {
	return c.bean.FormatCheck()
}