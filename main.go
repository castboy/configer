package main

import (
	"configer/server/base"
	"configer/server/structure"
	"fmt"
)

func main() {
	a := &structure.Symbol{Symbol: "AUDCAD"}
	exist, err := base.Get(base.NewSymboler(a))
	fmt.Println(a, exist, err)

	b := &structure.Symbol{Symbol: "AUDCAD"}
	num, err := base.Insert(base.NewSymboler(b))
	fmt.Println(num, err)
}
