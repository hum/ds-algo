package algorithms

func InsertionSort(values []int) []int {
	for i := 0; i < len(values); i++ {
		for j := 0; j < i+1; j++ {
			if values[j] > values[i] {
				values[j], values[i] = values[i], values[j]
			}
		}
	}
	return values
}
