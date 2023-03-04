package util

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := make([]int, 0)
	right := make([]int, 0)
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}
