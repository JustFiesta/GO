package main

import "fmt"

func main() {

	slice_i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(slice_i...))
	fmt.Println(bar(slice_i))
}

// multiple parameters of ONE type
func foo(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}

	return sum
}

// multiple return function
func bar(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}

	return sum
}
