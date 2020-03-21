package main

import (
	"configer/server"
	"configer/server/structure"
	"fmt"
)

func main() {
	a := &structure.A{Name: "wmq"}
	exist, err := server.Get(server.NewAer(a))
	fmt.Println(a, exist, err)

	b := &structure.A{Name: "wxx", Age: 30}
	num, err := server.Insert(server.NewAer(b))
	fmt.Println(num, err)
}
