package utils

import (
	"configer/server/structure"
)

func BuildConvInfo(currency string, sources []structure.Source) *structure.ConvInfo {
	if currency == "USD" {
		return &structure.ConvInfo{}
	}

	for j := range sources {
		cs := &structure.ConvInfo{
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
