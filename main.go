package main

import (
	"configer/server"
	"configer/server/structure"
	"fmt"
)

func main() {
	a := &structure.Symbol{Symbol: "AUDCAD"}
	exist, err := server.Get(server.NewSymboler(a))
	fmt.Println(a, exist, err)

	b := &structure.Symbol{Symbol: "AUDCAD"}
	num, err := server.Insert(server.NewSymboler(b))
	fmt.Println(num, err)
}
