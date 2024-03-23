package main

import "fmt"

func get42() int {
	return 42
}

func add42(num int, function_param func() int) int{

	sum := num + function_param()

	return sum
}

func main() {
	// func(name func(num int) float) string{}
	fmt.Println("Adding 42 to num = 10 is: ", add42(10, get42))
}
