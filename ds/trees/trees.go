package trees

import (
	"fmt"
)

type node struct {
	Value string
	Left  *node
	Right *node
}

func Preorder(root *node) { //visita raiz aintes
	if root != nil {
		fmt.Printf("%v", root.Value)
		Preorder(root.Left)
		Preorder(root.Right)
	}
}

func Postorder(root *node) { //visita raiz depois
	if root != nil {
		Postorder(root.Left)
		Postorder(root.Right)
		fmt.Printf("%v", root.Value)

	}
}

func Inorder(root *node) { //visita raiz entre
	if root != nil {
		Inorder(root.Left)
		fmt.Println(root)
		Inorder(root.Right)
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
		fmt.Println(node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

type arrayTree struct {
	tree []int
}

func NewArrayTree(val int) *arrayTree {
	treeArr := make([]int, 100)
	treeArr[0] = val
	return &arrayTree{
		tree: treeArr[:],
	}
}

func NewTree(value string) *node {
	return &node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func TransformToHeap(arr *[]int) *arrayTree {
	at := NewArrayTree(0)
	at.tree = *arr
	for start := len(at.tree) / 2; start >= 0; start-- {
		at.heapifyDown(start)
	}

	return at
}

func HeapSort(arr []int) *arrayTree {
	newHeap := TransformToHeap(&arr)
	for i := range len(newHeap.tree) {
		newHeap.heapifyDown(i)
	}
	return newHeap
}

func (at *arrayTree) Insert(val int) {
	at.tree = append(at.tree, val)
	at.heapifyUp(len(at.tree) - 1)
}

func (at *arrayTree) ExtractRoot() {
	lastIndex := len(at.tree) - 1
	at.tree[0] = at.tree[lastIndex]
	at.tree = at.tree[:lastIndex]
	at.heapifyDown(0)
}

func (at *arrayTree) heapifyUp(pos int) {
	if pos == 0 {
		return
	}
	if *at.parent(pos) < at.getValue(pos) {
		at.changeWithParent(pos)
		at.heapifyUp((pos - 1) / 2)
	}
}

func (at *arrayTree) parent(pos int) *int {
	return &at.tree[(pos-1)/2]
}

func (at *arrayTree) changeWithParent(pos int) {
	at.tree[pos], *at.parent(pos) = *at.parent(pos), at.tree[pos]

}

func (at *arrayTree) heapifyDown(pos int) {
	leftPos := (pos * 2) + 1
	rightPos := (pos * 2) + 2
	if leftPos > len(at.tree) || rightPos > len(at.tree) {
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

func (at *arrayTree) addLeft(pos, val int) {
	*at.left(pos) = val
}

func (at *arrayTree) addRight(pos, val int) {
	*at.right(pos) = val
}

func (at *arrayTree) left(pos int) *int {
	return &at.tree[(pos*2)+1]
}

func (at *arrayTree) right(pos int) *int {
	return &at.tree[(pos*2)+2]
}

func (at *arrayTree) getValue(pos int) int {
	return at.tree[pos]
}
