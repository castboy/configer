package server

// all functions under this package

func Insert(a configor) error {
	_, err := a.tableInsert()
	if err != nil {
		return err
	}

	return a.cacheInsert()
}

func Delete(a configor) error {
	num, err := a.tableDelete()
	if err != nil {
		return err
	}

	if num == 0 {
	}

	return a.cacheDelete()
}

func Update(a configor) error {
	num, err := a.tableUpdate()
	if err != nil {
		return err
	}

	if num == 0 {

	}

	return a.cacheUpdate()
}

func Get(a configor) error {
	hit := a.cacheGet()
	if hit {
		return nil
	}

	hit, err := a.tableGet()
	if err != nil {
		return err
	}

	if !hit {

	}

	return nil
}
