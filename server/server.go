package server

// all functions under this package

func Insert(a Configor) (num int64, err error) {
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

func Delete(a Configor) (num int64, err error) {
	num, err = a.GetTabler().Delete()
	if err != nil {
		return
	}

	return a.GetCacher().Delete()
}

func Update(a Configor) (num int64, err error) {
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

func Get(a Configor) (exist bool, err error) {
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

func Cache(a Configor) error {
	i, err := a.GetTabler().Export()
	if err != nil {
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}

func ExtendCache(a ExtendOperator) error {
	i, err := a.GetTabler().Export()
	if err != nil {
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}