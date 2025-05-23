package trees

import (
	"fmt"
)

type BinaryTree interface {
	Insert(val int)
	Left(pos int) int
	Right(pos int) int
	Parent(pos int) int
}

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

func BreadthFirst(root *Node) { // visita por alturas
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

type ArrayTree struct {
	tree []int
}

func NewArrayTree(val int) *ArrayTree {
	var treeArr [100]int
	return &ArrayTree{
		tree: treeArr[:],
	}
}

func (at *ArrayTree) Parent(pos int) *int {
	return &at.tree[(pos-1)/2]
}

func (at *ArrayTree) changeWithParent(pos int) {
	temp := at.Parent(pos)
	*at.Parent(pos) = at.GetValue(pos)
	at.tree[pos] = *temp
}

func (at *ArrayTree) HeapifyUp(pos int) {
	if *at.Parent(pos) < at.GetValue(pos) {
		at.changeWithParent(pos)
		at.HeapifyUp(pos)
	}
}

func (at *ArrayTree) AddLeft(pos, val int) {
	*at.Left(pos) = val
}

func (at *ArrayTree) AddRight(pos, val int) {
	*at.Right(pos) = val
}

func (at *ArrayTree) Left(pos int) *int {
	return &at.tree[(pos*2)+1]
}

func (at *ArrayTree) Right(pos int) *int {
	return &at.tree[(pos*2)+2]
}

func (at *ArrayTree) GetValue(pos int) int {
	return at.tree[pos]
}

func NewTree(val any) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}
