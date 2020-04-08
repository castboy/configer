package cache

type CacheOperator interface {
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	Get() (bool, error)
	Export() (i interface{}, err error)

	Cache(i interface{})
}

