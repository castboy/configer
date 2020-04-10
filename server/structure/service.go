package structure

type Checkor interface {
	FormatCheck() error
}

type Xormor interface {
	AutoCondition() string
}

