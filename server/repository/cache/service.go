package cache

type BaseOperator interface {
	Insert()
	Delete()
	Update()
	Get() (interface{}, bool)
	Export() (i interface{}, err error)

	ExtendOperator
}

type ExtendOperator interface {
	Cache(i interface{})
}
