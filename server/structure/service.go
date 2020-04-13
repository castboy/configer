package structure

type Checkor interface {
	FormatCheck() error
}

type Xormor interface {
	AutoCondition() string
}

type Cacheor interface {
	GetName() string
	GetID() int
}

