package indexNameID

import (
	"configer/server/structure/indexNameID"
)

type NameIDor interface {
	Insert(bean indexNameID.NameIDor)
	Delete(bean indexNameID.NameIDor)
	Update(bean indexNameID.NameIDor)
	Get(bean indexNameID.NameIDor) (indexNameID.NameIDor, bool)
	Export() (i interface{}, err error)
}
