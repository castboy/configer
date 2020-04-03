package server

// any structure can call functions under this package, if it implement interface below.

type tableOperator interface {
	baseOperator
}

type cacheOperator interface {
	baseOperator

	Cache(i interface{})
}

type checkor interface {
	FormatCheck() error
}

type baseOperator interface {
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	Get() (bool, error)
	Export() (i interface{}, err error)
}

type configor interface {
	GetTabler() tableOperator
	GetCacher() cacheOperator
	GetChecker() checkor
}

