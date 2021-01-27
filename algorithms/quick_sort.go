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
			swap(arr, i, partitionIndex)
			partitionIndex++
		}
	}

	swap(arr, right, partitionIndex)
	return partitionIndex
}

func swap(arr []int, first, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}
