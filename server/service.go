package server

// any structure can call functions under this package, if it implement interface below.

type configor interface {
	tableor
	cacheor
}

type tableor interface {
	tableInsert() (int64, error)
	tableDelete() (int64, error)
	tableUpdate() (int64, error)
	tableGet() (bool, error)
}

type cacheor interface {
	cacheInsert() error
	cacheDelete() error
	cacheUpdate() error
	cacheGet() bool
}
