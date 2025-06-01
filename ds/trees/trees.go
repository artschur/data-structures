package trees

import (
	"fmt"
)

type node struct {
	value string
	left  *node
	right *node
}

func NewTree(val string) *node {
	return &node{
		value: val,
		left:  nil,
		right: nil,
	}
}

func Preorder(root *node) { //visita raiz aintes
	if root != nil {
		fmt.Printf("%v", root.value)
		Preorder(root.left)
		Preorder(root.right)
	}
}

func Postorder(root *node) { //visita raiz depois
	if root != nil {
		Postorder(root.left)
		Postorder(root.right)
		fmt.Printf("%v", root.value)

	}
}

func Inorder(root *node) { //visita raiz entre
	if root != nil {
		Inorder(root.left)
		fmt.Println(root)
		Inorder(root.right)
	}
}

func BreadthFirst(root *node) { // visita por alturas
	if root == nil {
		return
	}
	queue := []*node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // remove the first element
		fmt.Println(node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
