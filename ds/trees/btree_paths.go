package trees

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bnode struct {
	value int
	left  *bnode
	right *bnode
}

func read_input_Btree() []int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	Nodes_arr := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		new_number, err := strconv.Atoi(line)
		if err != nil {
			panic("nonumber")
		}
		Nodes_arr = append(Nodes_arr, new_number)
	}

	if err := scanner.Err(); err != nil {
		panic("error reading input")
	}

	return Nodes_arr
}

func insert_element(root *bnode, element int) *bnode {
	if root == nil {
		return &bnode{
			value: element,
			left:  nil,
			right: nil,
		}
	}

	//till end
	switch {
	case element < root.value:
		root.left = insert_element(root.left, element)
	case element > root.value:
		root.right = insert_element(root.right, element)
	}
	return root
}

func preorderBST(root *bnode) { //visita raiz aintes
	if root != nil {
		fmt.Printf("%d ", root.value)
		preorderBST(root.left)
		preorderBST(root.right)
	}
}

func postorderBST(root *bnode) { //visita raiz depois
	if root != nil {
		postorderBST(root.left)
		postorderBST(root.right)
		fmt.Printf("%d ", root.value)

	}
}

func inorderBST(root *bnode) { //visita raiz entre
	if root != nil {
		inorderBST(root.left)
		fmt.Printf("%d ", root.value)
		inorderBST(root.right)
	}
}

func BinaryTreePathsSolution() {
	input := read_input_Btree()
	// fmt.Println(input)
	var root *bnode
	for _, el := range input {
		root = insert_element(root, el)
	}

	preorderBST(root)
	fmt.Println("")
	inorderBST(root)
	fmt.Println("")
	postorderBST(root)
}
