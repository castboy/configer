package cache

import "configer/server/structure"

type CacherA struct {
	a *structure.A
}

var cacherA *CacherA

func NewCacherA(a *structure.A) *CacherA {
	if cacherA == nil {
		cacherA = &CacherA{a}
	}

	return cacherA
}

// implement Cacheor
func (c *CacherA) Insert() (num int64, err error) {
	return
}

func (c *CacherA) Delete() (num int64, err error) {
	return
}

func (c *CacherA) Update() (num int64, err error) {
	return
}

func (c *CacherA) Get() (exist bool, err error) {
	return
}