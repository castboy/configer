package base

func Insert(a BaseOperator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	num, err = a.GetTabler().Insert()
	if err != nil {
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a BaseOperator) (num int64, err error) {
	num, err = a.GetTabler().Delete()
	if err != nil {
		return
	}

	return a.GetCacher().Delete()
}

func Update(a BaseOperator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	num, err = a.GetTabler().Update()
	if err != nil {
		return
	}

	return a.GetCacher().Update()
}

func Get(a BaseOperator) (exist bool, err error) {
	exist, _ = a.GetCacher().Get()
	if !exist {
		exist, err = a.GetTabler().Get()
		err = a.GetChecker().FormatCheck()
		if err != nil {
			exist = false
			return
		}

		if exist {
			a.GetCacher().Insert()
		}
	}

	return
}

func Export(a BaseOperator) (i interface{}, err error) {
	return a.GetCacher().Export()
}

func Cache(a BaseOperator) error {
	i, err := a.GetTabler().Export()
	if err != nil {
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}

