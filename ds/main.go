package main

import (
	"ed/trees"
	"fmt"
)

func main() {
	nt := trees.NewArrayTree(10)
	nt.AddLeft(1, 20)
	fmt.Println(nt)
}
