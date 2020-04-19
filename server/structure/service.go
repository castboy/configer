package structure

type Helpor interface {
	Checkor
	NotFoundError() error
}

type Checkor interface {
	FormatCheck() error
	IndexCheck() error
}

type Xormor interface {
	AutoCondition() string
}
