package trees

// all nodes need to be from -1 to 1 to be balanced tree
// check on insert
type balanced_node struct {
	key         int
	value       int
	left        *balanced_node
	right       *balanced_node
	load_factor int // h right - h_left.
	height      int
}

func NewBalancedTree() *balanced_node {
	return &balanced_node{
		key:         0,
		value:       0,
		left:        nil,
		right:       nil,
		load_factor: 0,
		height:      1,
	}
}

func (n *balanced_node) getLoadFactor() int {
	if n == nil {
		return 0
	}
	leftHeight := n.left.getHeight()
	rightHeight := n.right.getHeight()
	return rightHeight - leftHeight
}

func (n *balanced_node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *balanced_node) updateHeight() {
	if n == nil {
		return
	}

	leftHeight := n.left.getHeight()
	rightHeight := n.right.getHeight()

	if leftHeight > rightHeight {
		n.height = leftHeight + 1
	} else {
		n.height = rightHeight + 1
	}
}

func (n *balanced_node) rotateRight() *balanced_node {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

func (n *balanced_node) rotateLeft() *balanced_node {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	// Update heights
	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

func (n *balanced_node) Insert(val int) *balanced_node {
	if n == nil {
		return &balanced_node{
			key:         val,
			value:       val,
			left:        nil,
			right:       nil,
			load_factor: 0,
			height:      1,
		}
	}

	if val > n.value {
		n.right = n.right.Insert(val)
	} else {
		n.left = n.left.Insert(val)
	}
	n.updateHeight()

}

func main() {
	// tree := NewBalancedTree()
}
