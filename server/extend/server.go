package extend

import "configer/server/constant"

// all functions under this package

func Insert(a operator) (err error) {
	err = a.GetHelper().FormatCheck()
	if err != nil {
		return
	}

	a.GetCacher().Insert()

	return nil
}

func Delete(a operator) {
	a.GetCacher().Delete()
}

func Update(a operator) error {
	err := a.GetHelper().FormatCheck()
	if err != nil {
		return err
	}

	a.GetCacher().Update()

	return nil
}

func Get(a operator) (i interface{}, err error) {
	err = a.GetHelper().IndexCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	i, exist := a.GetCacher().Get()
	if !exist {
		err = constant.NewErr(constant.ArgsErr, a.GetHelper().NotFoundError())
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
