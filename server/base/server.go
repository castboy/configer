package base

import (
	"configer/server/constant"
	"fmt"
)

func Insert(a operator) (err error) {
	err = a.GetHelper().FormatCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	num, err := a.GetTabler().Insert()
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return
	}

	if num != 1 {
		err = constant.NewErr(constant.UnknowErr, fmt.Sprintf("Insert 0 item, no error, structure: %v", a.GetHelper()))
		return
	}

	a.GetCacher().Insert()

	return nil
}

func Delete(a operator) (num int64, err error) {
	num, err = a.GetTabler().Delete()
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return
	}

	a.GetCacher().Delete()

	return
}

func Update(a operator) (num int64, err error) {
	err = a.GetHelper().FormatCheck()
	if err != nil {
		err = constant.NewErr(constant.ArgsErr, err)
		return
	}

	num, err = a.GetTabler().Update()
	fmt.Println("num",num)
	if err != nil {
		err = constant.NewErr(constant.DBErr, err)
		return
	}

	if num == 0 {
		return
	}

	a.GetCacher().Update()

	return
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
		err = constant.NewErr(constant.DBErr, err)
		return err
	}

	a.GetCacher().Cache(i)

	return nil
}
