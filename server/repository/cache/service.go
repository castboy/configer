package cache

type BaseOperator interface {
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	Get() (interface{}, bool)
	Export() (i interface{}, err error)

	ExtendOperator
}

type ExtendOperator interface {
	Cache(i interface{})
}