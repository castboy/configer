package structure

type FullSymbolName struct {
	Sl       SymbolLeverage
	FullName string
}

type SymbolLeverage struct {
	Symbol   string
	Leverage int32
}

func (fs *FullSymbolName) FormatCheck() error {
	return nil
}

func (fs *FullSymbolName) IndexCheck() error {
	return nil
}

func (fs *FullSymbolName) AutoCondition() (cond string) {
	return
}
