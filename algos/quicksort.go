package algos

// worst case is when its already ordered.
func Quick_sort(arr []int, start, end int) []int {
	if start < end {
		q := partition(arr, start, end)
		Quick_sort(arr, start, q-1)
		Quick_sort(arr, q+1, end)
	}
	return arr
}

//ideal eh o particionamento dividir igualmente

func partition(arr []int, start, end int) (midIndex int) {
	pivot := arr[end]
	i := start - 1
	for it_idx := start; it_idx < end; it_idx++ {
		if arr[it_idx] <= pivot {
			i += 1
			arr[i], arr[it_idx] = arr[it_idx], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1 //return pivot pos

}
