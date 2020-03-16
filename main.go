package main

import (
	"configer/server"
	"configer/server/structure"
	"fmt"
)

func main() {
	a := &structure.A{Name: "wmq"}
	aer := server.NewAer(a)
	err := server.Get(aer)
	if err == nil {
		fmt.Println(err)
	}

	fmt.Println(a)
}
