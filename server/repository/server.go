package repository


func Insert(a BaseOperator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a BaseOperator) (num int64, err error) {
	return a.GetCacher().Delete()
}

func Update(a BaseOperator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Update()
}

func Get(a BaseOperator) (exist bool, err error) {
	return a.GetCacher().Get()
}

func Export(a BaseOperator) (i interface{}, err error) {
	return a.GetCacher().Export()
}