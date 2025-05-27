package queue

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([8][8]int, state) {
	var matrix [8][8]int
	var startingPos state

	scanner := bufio.NewScanner(os.Stdin)
	letterMapping := map[string]int{
		"a": 0, "b": 1, "c": 2, "d": 3, "e": 4, "f": 5,
		"g": 6, "h": 7,
	}

	var isFinal bool = false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		row, ok := letterMapping[string(line[0])]

		if !ok {
			fmt.Fprintln(os.Stderr, "Invalid row letter:", line[0])
			continue
		}

		col, err := strconv.Atoi(string(line[1]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid column number:", line[1])
			continue
		}

		col--
		if isFinal {
			matrix[row][col] = 2
		} else {
			matrix[row][col] = 1 //cavalo
			startingPos = state{
				Coord: coord{Row: row, Col: col}, Distance: 0,
			}
		}
		isFinal = !isFinal
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
	return matrix, startingPos
}

func findShortestPath(matrix [8][8]int, posInit state) int {
	q := Queue{}
	q.Push(posInit)

	visited := [8][8]bool{}
	visited[posInit.Coord.Row][posInit.Coord.Col] = true

	for !q.IsEmpty() {
		current := q.Pop()
		if matrix[current.Coord.Row][current.Coord.Col] == 2 {
			return current.Distance
		}
		var knightMoves = [][2]int{
			{-2, -1}, {-2, 1}, {-1, -2},
			{-1, 2}, {1, -2}, {1, 2}, {2, -1},
			{2, 1},
		}

		for _, move := range knightMoves {
			nextRow, nextCol := current.Coord.Row+move[0], current.Coord.Col+move[1]

			if nextRow >= 0 && nextRow < 8 && nextCol >= 0 && nextCol < 8 && !visited[nextRow][nextCol] {
				visited[nextRow][nextCol] = true
				nextState := state{
					Coord:    coord{Row: nextRow, Col: nextCol},
					Distance: current.Distance + 1,
					// atual + 1
				}
				q.Push(nextState)
			}
		}
	}
	return -1
}

func MinimumKnightMoves() {
	matrix, startingPos := parseInput()

	shortestPath := findShortestPath(matrix, startingPos)
	fmt.Println("Movimentos: ", shortestPath)
}
