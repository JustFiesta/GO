package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
	fmt.Println(bar())
}

// basic function
func foo() int {
	return 42
}

// multiple return function
func bar() (int, string) {
	return 42, "hello"
}
