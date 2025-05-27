package main

import (
	"ed/trees"
	"fmt"
)

func main() {
	nt := trees.NewArrayTree(10)
	nt.ExtractRoot()
	fmt.Println(nt)
}
