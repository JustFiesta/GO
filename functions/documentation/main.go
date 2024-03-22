package main

import "fmt"

// Add two integers
func Add(a, b int) int {
	return a + b
}

// Multiply two integers
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Add(5, 7))
}
