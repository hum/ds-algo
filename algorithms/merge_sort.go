package algorithms

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0)

	for len(arr1) > 0 || len(arr2) > 0 {
		if len(arr1) == 0 {
			return append(result, arr2...)
		} else if len(arr2) == 0 {
			return append(result, arr1...)
		}

		if arr1[0] <= arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}
	return result
}

func MergeSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	pivot := len(values) / 2
	left := MergeSort(values[:pivot])
	right := MergeSort(values[pivot:])

	return merge(left, right)
}
