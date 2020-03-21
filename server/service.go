package server

// any structure can call functions under this package, if it implement interface below.

type operator interface {
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	Get() (bool, error)
}

type configor interface {
	GetCacher() operator
	GetTabler() operator
}