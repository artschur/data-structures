package main

import (
	"ed/trees"
)

func main() {
	input := trees.ReadInput()
	root := trees.Mount_tree(input)
	trees.Preorder(root)
	trees.Postorder(root)

}
