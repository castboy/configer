package nameIDor

import (
	"configer/server/structure"
)

type NameIDor interface {
	Insert(bean structure.NameIDor)
	Delete(bean structure.NameIDor)
	Update(bean structure.NameIDor)
	Get(bean structure.NameIDor) (structure.NameIDor, bool)
	Export() (i interface{}, err error)
}
