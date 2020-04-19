package extend

import "configer/server/constant"

// all functions under this package

func Insert(a operator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a operator) (num int64, err error) {
	return a.GetCacher().Delete()
}

func Update(a operator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		return
	}

	return a.GetCacher().Update()
}

func Get(a operator) (i interface{}, err error) {
	err = a.GetChecker().IndexCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	i, exist := a.GetCacher().Get()
	if !exist {
		err = constant.NewErr(constant.ArgsErr, a.GetChecker().NotFoundError())
		return
	}

	return
}

func Export(a operator) (i interface{}, err error) {
	return a.GetCacher().Export()
}

func Cache(a operator) error {
	i, err := a.GetTabler().Export()
	if err != nil {
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}
