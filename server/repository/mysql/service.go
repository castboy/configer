package mysql

type BaseOperator interface {
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	Get() (bool, error)

	ExtendOperator
}

type ExtendOperator interface {
	Export() (i interface{}, err error)
}
