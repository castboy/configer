package indexName

import (
	"configer/server/structure/indexName"
)

type Nameor interface {
	Insert(bean indexName.Nameor)
	Delete(bean indexName.Nameor)
	Update(bean indexName.Nameor)
	Get(bean indexName.Nameor) (indexName.Nameor, bool)
	Export() (i interface{}, err error)
}
