package cache

import "configer/server/structure"

type CacherA struct {

}

var cacherA *CacherA

func GetCacherA() *CacherA {
	if cacherA == nil {
		cacherA = &CacherA{}
	}

	return cacherA
}

// implement Cacheor
func (c *CacherA) Insert(a *structure.A) error {
	return nil
}

func (c *CacherA) Delete(a *structure.A) error {
	return nil
}

func (c *CacherA) Update(a *structure.A) error {
	return nil
}

func (c *CacherA) Get(a *structure.A) bool {
	return false
}