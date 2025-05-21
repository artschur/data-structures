package main

import (
	"ed/trees"
)

func main() {
	root := trees.NewTree(10)
	root.Left = trees.NewTree(5)
	root.Right = trees.NewTree(15)
	root.Left.Left = trees.NewTree(3)
	root.Left.Right = trees.NewTree(7)
	root.Right.Left = trees.NewTree(12)
	root.Right.Right = trees.NewTree(18)
	trees.Preorder(root)
	trees.Postorder(root)

}
