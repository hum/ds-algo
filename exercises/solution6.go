package main

import (
	"fmt"
)

func findFactorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * findFactorialRecursive(value-1)
}

func findFactorialIterative(value int) int {
	if value == 2 {
		return 2
	}

	result := 1
	for i := value; i > 1; i-- {
		result *= i
	}
	return result
}

func main() {
	result := findFactorialRecursive(5)
	result2 := findFactorialIterative(5)
	fmt.Println(result, result2)
}
