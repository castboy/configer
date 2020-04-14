package indexID

import (
	"configer/server/structure/indexID"
)

type IDor interface {
	Insert(bean indexID.IDor)
	Delete(bean indexID.IDor)
	Update(bean indexID.IDor)
	Get(bean indexID.IDor) (indexID.IDor, bool)
	Export() (i interface{}, err error)
}
