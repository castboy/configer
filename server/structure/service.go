package structure

type Helpor interface {
	Checkor
	NotFoundError() error
}

type Checkor interface {
	FormatCheck() error
	IndexCheck() error
}

type Xormor interface {
	ExportCondition() string
	UpdateCondition() string
	DeleteCondition() string
	GetCondition() string
}
