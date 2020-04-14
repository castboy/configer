package indexID

import (
	"configer/server/structure"
)

type IDor interface {
	Insert(bean structure.IDor)
	Delete(bean structure.IDor)
	Update(bean structure.IDor)
	Get(bean structure.IDor) (structure.IDor, bool)
	Export() (i interface{}, err error)
}
