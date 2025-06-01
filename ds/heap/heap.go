package heap

// its a bin tree where the child is
// less than the root and its always aligned to the left
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

func TransformToHeap(arr *[]int) *ArrayTree {
	at := NewArrayTree(0)
	at.tree = *arr
	for start := len(at.tree) / 2; start >= 0; start-- {
		at.heapifyDown(start)
	}

	return at
}

func HeapSort(arr []int) *ArrayTree {
	newHeap := TransformToHeap(&arr)
	for i := range len(newHeap.tree) {
		newHeap.heapifyDown(i)
	}
	return newHeap
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

func (at *ArrayTree) parent(pos int) *int {
	return &at.tree[(pos-1)/2]
}

func (at *ArrayTree) changeWithParent(pos int) {
	at.tree[pos], *at.parent(pos) = *at.parent(pos), at.tree[pos]

}

func (at *ArrayTree) heapifyDown(pos int) {
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

func (at *ArrayTree) left(pos int) *int {
	return &at.tree[(pos*2)+1]
}

func (at *ArrayTree) right(pos int) *int {
	return &at.tree[(pos*2)+2]
}

func (at *ArrayTree) getValue(pos int) int {
	return at.tree[pos]
}
