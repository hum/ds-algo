package main

import (
	"fmt"
)

func reverseIterative(text string) string {
	result := ""
	data := []rune(text)

	for i := len(data) - 1; i >= 0; i-- {
		result += string(data[i])
	}
	return result
}

func reverseRecursive(text string) string {
	if len(text) <= 1 {
		return text
	}
	data := []rune(text)
	return reverseRecursive(string(data[1:])) + string(data[0])
}

func main() {
	text := "Hello, my name is..."
	result := reverseIterative(text)
	result2 := reverseRecursive(text)
	fmt.Println(result, result2)
}
