package main

import "fmt"

// add two integers
func add(a, b int) int {
	return a + b
}

// multiply two integers
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(add(5, 7))
}
