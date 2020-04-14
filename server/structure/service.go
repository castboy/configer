package structure

type Checkor interface {
	FormatCheck() error
	IndexCheck() error
}

type Xormor interface {
	AutoCondition() string
}

type IDor interface {
	GetID() int
}

type Nameor interface {
	GetName() string
}

type NameIDor interface {
	GetName() string
	GetID() int
}
