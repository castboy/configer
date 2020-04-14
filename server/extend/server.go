package extend

// all functions under this package

func Insert(a ExtendOperator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a ExtendOperator) (num int64, err error) {
	return a.GetCacher().Delete()
}

func Update(a ExtendOperator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Update()
}

func Get(a ExtendOperator) (exist bool, err error) {
	return a.GetCacher().Get()
}

func Export(a ExtendOperator) (i interface{}, err error) {
	return a.GetCacher().Export()
}

func Cache(a ExtendOperator) error {
	i, err := a.GetTabler().Export()
	if err != nil {
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}
