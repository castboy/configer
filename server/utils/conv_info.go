package utils

import (
	"configer/server/structure/indexName"
	"configer/server/structure/indexNameID"
)

func BuildConvInfo(currency string, sources []indexNameID.Source) *indexName.ConvInfo {
	if currency == "USD" {
		return &indexName.ConvInfo{}
	}

	for j := range sources {
		cs := &indexName.ConvInfo{
			ConvNeed:   true,
			ConvSymbol: sources[j].Source,
		}
		if sources[j].Source == currency+"USD" {
			cs.ConvMultiply = true
			return cs
		}
		if sources[j].Source == "USD"+currency {
			return cs
		}
	}

	return nil
}
