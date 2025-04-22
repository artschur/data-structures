package main

import (
	"bufio"
	"fmt"
	"largest-island/stack"
	"os"
	"strings"
)

func visitedMatrix(rows int, cols int) [][]bool {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	return visited
}

func parseInput() [][]byte {
	var matrix [][]byte
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // 👈 Trim spaces!
		row := []byte(line)
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	return matrix
}

func searchIsland(matrix [][]byte, visited [][]bool, startRow int, startCol int) int {
	if startRow < 0 || startCol < 0 || startRow >= len(matrix) || startCol >= len(matrix[0]) || visited[startRow][startCol] || matrix[startRow][startCol] == '0' {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])

	s := stack.NewStack()
	s.Push(stack.Coord{X: startRow, Y: startCol})
	visited[startRow][startCol] = true
	var size int = 0

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for s.Len > 0 {
		current := s.Peek()
		s.Pop()
		size++

		for _, dir := range directions {
			nextRow := current.X + dir[0]
			nextCol := current.Y + dir[1]
			//if has not been visited and is land
			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && !visited[nextRow][nextCol] && matrix[nextRow][nextCol] == '1' {
				s.Push(stack.Coord{X: nextRow, Y: nextCol})
				visited[nextRow][nextCol] = true
			}
		}
	}
	return size
}

func main() {

	matrix := parseInput()

	if len(matrix) == 0 {
		return
	}

	rows := len(matrix)
	cols := len(matrix[0])
	visited := visitedMatrix(rows, cols)
	maxSize := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' && !visited[i][j] {
				size := searchIsland(matrix, visited, i, j)
				if size > maxSize {
					maxSize = size
				}
			}
		}
	}
	fmt.Println("Max size of island:", maxSize)
}
