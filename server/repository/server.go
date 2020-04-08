package repository


func Insert(a Configor) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a Configor) (num int64, err error) {
	return a.GetCacher().Delete()
}

func Update(a Configor) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Update()
}

func Get(a Configor) (exist bool, err error) {
	return a.GetCacher().Get()
}

func Export(a Configor) (i interface{}, err error) {
	return a.GetCacher().Export()
}