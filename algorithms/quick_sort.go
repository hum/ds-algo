package algorithms

func QuickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	pivot := right
	partitionIndex := partition(arr, pivot, left, right)

	QuickSort(arr, left, partitionIndex-1)
	QuickSort(arr, partitionIndex+1, right)
	return arr
}

func partition(arr []int, pivot, left, right int) int {
	val := arr[pivot]
	partitionIndex := left

	for i := left; i < right; i++ {
		if arr[i] < val {
			arr[i], arr[partitionIndex] = arr[partitionIndex], arr[i]
			partitionIndex++
		}
	}

	arr[right], arr[partitionIndex] = arr[partitionIndex], arr[right]
	return partitionIndex
}
