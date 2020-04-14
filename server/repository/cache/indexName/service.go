package indexName

import (
	"configer/server/structure"
)

type Nameor interface {
	Insert(bean structure.Nameor)
	Delete(bean structure.Nameor)
	Update(bean structure.Nameor)
	Get(bean structure.Nameor) (structure.Nameor, bool)
	Export() (i interface{}, err error)
}
