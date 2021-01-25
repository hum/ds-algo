package main

import (
	"fmt"
)

func fibonacciIterative(value int) int {
	a, b := 0, 1

	for i := 0; i < value; i++ {
		a, b = b, a+b
	}
	return a
}

func fibonacciRecursive(value int) int {
	if value < 2 {
		return value
	}
	return fibonacciRecursive(value-1) + fibonacciRecursive(value-2)
}

func main() {
	result := fibonacciIterative(8)
	result2 := fibonacciRecursive(8)
	fmt.Println(result, result2)
}
