package trees

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func find_operator(expression string) int {
	low_priority := []string{"-", "+"}
	high_priority := []string{"*", "/"}

	for i := len(expression) - 1; i >= 0; i-- {
		char := string(expression[i])
		if slices.Contains(low_priority, char) {
			return i
		}
	}

	for i := len(expression) - 1; i >= 0; i-- {
		char := string(expression[i])
		if slices.Contains(high_priority, char) {
			return i
		}
	}

	return -1
}

func Mount_tree(expression string) *Node {
	if len(expression) <= 1 {
		return NewTree(expression)
	}

	operator_index := find_operator(expression)

	if operator_index == -1 {
		return NewTree(expression) //case its number
	}

	left_expression, right_expression := expression[:operator_index], expression[operator_index+1:]

	operator := string(expression[operator_index])

	root := NewTree(operator)

	root.Left = Mount_tree(left_expression)
	root.Right = Mount_tree(right_expression)

	return root
}

func main() {
	input := ReadInput()
	root := Mount_tree(input)
	Preorder(root)
	Postorder(root)
}
