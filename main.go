package main

import (
	"ed/algos"
	"math/rand"
	"time"
)

func GenerateRandomArray(size int) []int {
	// Seed the random number generator with current time
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		// Generate random numbers between 1 and 1000000
		arr[i] = rand.Intn(1000000) + 1
	}
	return arr
}

func main() {
	// input := trees.ReadInput()
	// root := trees.Mount_tree(input)
	// trees.Preorder(root)
	// trees.Postorder(root)
	//
	largeArr := GenerateRandomArray(1000000)
	// arr := []int{2, 8, 7, 1, 3, 5, 6, 4}
	largeArr = algos.Quick_sort(largeArr, 0, len(largeArr)-1)
	// fmt.Printf("largeArr: %v\n", largeArr)

}
