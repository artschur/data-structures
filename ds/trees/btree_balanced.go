package trees

type balanced_btree interface {
	rotate_right(node balanced_node)
	rotate_left(node balanced_node)
	insert(val int8)
}

// all nodes need to be from -1 to 1 to be balanced tree
// check on insert
type balanced_node struct {
	bnode
	load_factor int8 // h right - h_left.
	height      int8
}

func main() {

}
