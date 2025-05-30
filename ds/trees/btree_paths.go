package trees

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Read_input() []int8 {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	nodes_arr := []int8{}

	for scanner.Scan() {
		line := scanner.Text()
		new_number, err := strconv.Atoi(line)
		if err != nil {
			panic("nonumber")
		}
		nodes_arr = append(nodes_arr, int8(new_number))
	}

	if err := scanner.Err(); err != nil {
		panic("error reading input")
	}

	return nodes_arr
}

type Bnode struct {
	value int8
	left  *Bnode
	right *Bnode
}

func printSlice(slice []int8) {
	for i, val := range slice {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func Insert_element(root *Bnode, element int8) *Bnode {
	if root == nil {
		return &Bnode{value: element, left: nil, right: nil}
	}
	//vai progredindo, se chegar no final insere
	if element < root.value {
		root.left = Insert_element(root.left, element)
	} else {
		root.right = Insert_element(root.right, element)
	}
	return root
}

func InOrder(node *Bnode, result *[]int8) {
	if node != nil {
		InOrder(node.left, result)
		*result = append(*result, node.value)
		InOrder(node.right, result)
	}
}

func PostOrder(node *Bnode, result *[]int8) {
	if node != nil {
		PostOrder(node.left, result)
		PostOrder(node.left, result)
		*result = append(*result, node.value)
	}
}

func PreOrder(node *Bnode, result *[]int8) {
	if node != nil {
		*result = append(*result, node.value)
		PreOrder(node.left, result)
		PreOrder(node.right, result)

	}
}

func main() {
	input := Read_input()

	var root *Bnode
	for _, el := range input {
		root = Insert_element(root, el)
	}

	var preResult []int8
	PreOrder(root, &preResult)
	printSlice(preResult)

	var inResult []int8
	InOrder(root, &inResult)
	fmt.Printf("%v", inResult)
	printSlice(inResult)

	var postResult []int8
	PostOrder(root, &postResult)
	fmt.Printf("%v", postResult)
	printSlice(postResult)
}
