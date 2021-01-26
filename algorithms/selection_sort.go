package algorithms

func SelectionSort(values []int) []int {
	for i := 0; i < len(values); i++ {
		min := i

		for j := i+1; j < len(values); j++ {
			if values[j] < values[min] {
				min = j 
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}
