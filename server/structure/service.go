package structure

type Checkor interface {
	FormatCheck() error
}

type Xormor interface {
	AutoCondition() string
}

type Cacheor1 interface {
	GetName() string
	GetID() int
}

