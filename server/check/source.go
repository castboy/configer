package check

import "configer/server/structure"

type CheckerSource struct {
	bean *structure.Source
}

var checkerSource *CheckerSource

func NewCheckerSource(bean *structure.Source) *CheckerSource {
	if checkerSource == nil {
		checkerSource = &CheckerSource{
			bean,
		}
	}

	return checkerSource
}

func (c *CheckerSource) FormatCheck() error {
	return c.bean.FormatCheck()
}