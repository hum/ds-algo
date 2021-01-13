package main

import (
	"fmt"
	"go/types"
)

func main() {
	result, err := findRecurringNumber([]int{2, 5, 1, 2, 3, 5, 1, 2, 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func findRecurringNumber(arr []int) (int, error) {
	tmp := make(map[int]types.Nil, len(arr))

	for _, value := range arr {
		if _, ok := tmp[value]; ok {
			return value, nil
		}
		tmp[value] = types.Nil{}
	}
	return 0, fmt.Errorf("No pair of values found.")
}
