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
	treeArr := make([]int, 100)
	treeArr[0] = val
	return &ArrayTree{
		tree: treeArr[:],
	}
}

func (at *ArrayTree) parent(pos int) *int {
	return &at.tree[(pos-1)/2]
}

func (at *ArrayTree) changeWithParent(pos int) {
	at.tree[pos], *at.parent(pos) = *at.parent(pos), at.tree[pos]

}

func (at *ArrayTree) Insert(val int) {
	at.tree = append(at.tree, val)
	at.heapifyUp(len(at.tree) - 1)
}

func (at *ArrayTree) ExtractRoot() {
	lastIndex := len(at.tree) - 1
	at.tree[0] = at.tree[lastIndex]
	at.tree = at.tree[:lastIndex]
	at.heapifyDown(0)
}

func (at *ArrayTree) heapifyUp(pos int) {
	if pos == 0 {
		return
	}
	if *at.parent(pos) < at.getValue(pos) {
		at.changeWithParent(pos)
		at.heapifyUp((pos - 1) / 2)
	}
}

func (at *ArrayTree) heapifyDown(pos int) {
	leftPos := (pos * 2) + 1
	rightPos := (pos * 2) + 2
	if leftPos < len(at.tree) || rightPos < len(at.tree) {
		return
	}

	largest := pos
	if leftPos < len(at.tree) && *at.left(pos) > at.getValue(pos) {
		largest = leftPos
	}

	if rightPos < len(at.tree) && *at.right(pos) > at.getValue(pos) {
		largest = rightPos
	}

	if largest != pos {
		at.changeWithParent(largest)
		at.heapifyDown(largest)
	}
}

func (at *ArrayTree) addLeft(pos, val int) {
	*at.left(pos) = val
}

func (at *ArrayTree) addRight(pos, val int) {
	*at.right(pos) = val
}

func (at *ArrayTree) left(pos int) *int {
	return &at.tree[(pos*2)+1]
}

func (at *ArrayTree) right(pos int) *int {
	return &at.tree[(pos*2)+2]
}

func (at *ArrayTree) getValue(pos int) int {
	return at.tree[pos]
}

func NewTree(val any) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}
