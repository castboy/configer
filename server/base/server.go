package base

import "configer/server/constant"

func Insert(a operator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	num, err = a.GetTabler().Insert()
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return
	}

	return a.GetCacher().Insert()
}

func Delete(a operator) (num int64, err error) {
	num, err = a.GetTabler().Delete()
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return
	}

	return a.GetCacher().Delete()
}

func Update(a operator) (num int64, err error) {
	err = a.GetChecker().FormatCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	num, err = a.GetTabler().Update()
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return
	}

	return a.GetCacher().Update()
}

func Get(a operator) (i interface{}, exist bool, err error) {
	err = a.GetChecker().IndexCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	i, exist = a.GetCacher().Get()

	return
}

func Export(a operator) (i interface{}, err error) {
	return a.GetCacher().Export()
}

func Cache(a operator) error {
	i, err := a.GetTabler().Export()
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}
