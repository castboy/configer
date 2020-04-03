package mysql

import "configer/server/repository"

type TableOperator interface {
	repository.BaseOperator
}

