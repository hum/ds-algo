package main

import "fmt"

func main() {
	result := mergeSortedArrays([]int{0, 3, 4, 31}, []int{4, 6, 30})
	fmt.Println(result)
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	index1, index2 := 0, 0

	for {
		if index1 == len(arr1) {
			result = append(result, arr2[index2:]...)
			break
		}

		if index2 == len(arr2) {
			result = append(result, arr1[index1:]...)
			break
		}

		value := 0
		element1, element2 := arr1[index1], arr2[index2]

		if element1 < element2 {
			value = element1
			index1++
		} else {
			value = element2
			index2++
		}

		result = append(result, value)
	}
	return result
}
