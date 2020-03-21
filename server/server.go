package server

// all functions under this package

func Insert(a configor) (num int64, err error) {
	num, err = a.GetTabler().Insert()
	if err != nil {
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a configor) (num int64, err error) {
	num, err = a.GetCacher().Delete()
	if err != nil {
		return
	}

	return a.GetCacher().Delete()
}

func Update(a configor) (num int64, err error) {
	num, err = a.GetCacher().Update()
	if err != nil {
		return
	}

	return a.GetCacher().Update()
}

func Get(a configor) (exist bool, err error) {
	exist, _ = a.GetCacher().Get()
	if !exist || err != nil {
		exist, err = a.GetTabler().Get()
	}

	return
}
