package structure

type ConvSymbol struct {
	ConvInfo
	ConvType
	SourceName string
}

type ConvInfo struct {
	ConvNeed,
	ConvMultiply bool
	ConvSymbol string
}

type ConvType int

const (
	MarginConv ConvType = iota
	ProfitConv
)

func (cs *ConvSymbol) FormatCheck() error {
	return nil
}