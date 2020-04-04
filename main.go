package main

import (
	"configer/server"
	"configer/server/implement"
	"configer/server/structure"
	"fmt"
)

func main() {
	a := &structure.Symbol{Symbol: "AUDCAD"}
	exist, err := server.Get(implement.NewSymboler(a))
	fmt.Println(a, exist, err)

	b := &structure.Symbol{Symbol: "AUDCAD"}
	num, err := server.Insert(implement.NewSymboler(b))
	fmt.Println(num, err)
}
