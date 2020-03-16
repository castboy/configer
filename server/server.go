package server

type Servor interface {
	Tableor
	Cacheor
}

type Tableor interface {
	TableInsert() (int64, error)
	TableDelete() (int64, error)
	TableUpdate() (int64, error)
	TableGet() (bool, error)
}

type Cacheor interface {
	CacheInsert() error
	CacheDelete() error
	CacheUpdate() error
	CacheGet() bool
}

func Insert(a Servor) error {
	_, err := a.TableInsert()
	if err != nil {
		return err
	}

	return a.CacheInsert()
}

func Delete(a Servor) error {
	num, err := a.TableDelete()
	if err != nil {
		return err
	}

	if num == 0 {
	}

	return a.CacheDelete()
}

func Update(a Servor) error {
	num, err := a.TableUpdate()
	if err != nil {
		return err
	}

	if num == 0 {

	}

	return a.CacheUpdate()
}

func Get(a Servor) error {
	hit := a.CacheGet()
	if hit {
		return nil
	}

	hit, err := a.TableGet()
	if err != nil {
		return err
	}

	if !hit {

	}

	return nil
}
