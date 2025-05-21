package trees

import "fmt"

type Node struct {
	Value any
	Left  *Node
	Right *Node
}

func Preorder(root *Node) { //visita raiz aintes
	if root != nil {
		fmt.Println(root.Value)
		Preorder(root.Left)
		Preorder(root.Right)
	}
}

func Postorder(root *Node) { //visita raiz depois
	if root != nil {
		Postorder(root.Left)
		Postorder(root.Right)
		fmt.Println(root.Value)
	}
}

func Inorder(root *Node) { //visita raiz entre
	if root != nil {
		Inorder(root.Left)
		fmt.Println(root)
		Inorder(root.Right)
	}
}

func BreadthFirst(root *Node) { //visita raiz entre
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // remove the first element
		fmt.Println(node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func NewTree(val any) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}
