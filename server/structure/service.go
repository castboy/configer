package structure

type Checkor interface {
	FormatCheck() error
	IndexCheck() error
}

type Xormor interface {
	AutoCondition() string
}
