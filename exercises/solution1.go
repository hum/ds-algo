package main

import (
	"fmt"
)

func main() {
	result := reverse("Hello, my name is...")
	fmt.Println(result)
}

func reverse(s string) string {
	length := len(s)
	runes := make([]rune, length)

	for _, v := range s {
		length--
		runes[length] = v
	}
	return string(runes)
}
