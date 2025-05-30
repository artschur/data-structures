package main

import (
	"ed/trees"
	"fmt"
)

func main() {
	input := trees.Read_input()

	var root *trees.Bnode
	for _, el := range input {
		root = trees.Insert_element(root, el)
	}

	var inResult []int8
	fmt.Println()
	trees.InOrder(root, &inResult)
}
