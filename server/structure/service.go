package structure

type Checkor interface {
	FormatCheck() error
	IndexCheck() error

	NotFoundError() error
}

type Xormor interface {
	AutoCondition() string
}
