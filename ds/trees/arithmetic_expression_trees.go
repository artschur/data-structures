package trees

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func readInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
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

func mount_tree(expression string) *Node {
	if len(expression) <= 1 || isNumber(expression) {
		return NewTree(expression)
	}

	operator_index := find_operator(expression)

	if operator_index == -1 {
		return NewTree(expression) //case its number
	}

	left_expression, right_expression := expression[:operator_index], expression[operator_index+1:]

	operator := string(expression[operator_index])

	root := NewTree(operator)

	root.Left = mount_tree(left_expression)
	root.Right = mount_tree(right_expression)

	return root
}

func main() {
	input, err := readInput()
	strings.ReplaceAll(input, " ", "")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	mount_tree("")
}
