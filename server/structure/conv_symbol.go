package structure

type ConvSymbol struct {
	*ConvInfo
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
	ConvTypeLength
)

func (cs *ConvSymbol) FormatCheck() error {
	return nil
}

func (cs *ConvSymbol) AutoCondition() (cond string) {
	return
}

func (cs *ConvSymbol) GetName() string {
	return cs.SourceName
}