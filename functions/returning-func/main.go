package main

import "fmt"

func hello() func() int {

	fmt.Println("Hello")

	return get42
}

func get42() int {
	return 42
}

// 2nd way

func outer() func() int {
	return func() int {
		return 42
	}
}

func main() {
	// Func return enables code to return other function and run it

	x := hello()
	y := outer()

	fmt.Println("Called return function from x variable: ", x())
	fmt.Println("Called return function from y variable: ", y())

	// whithout the () after assigning func expression - it would print memory addres of function (bcs functrion is also a type!)
	fmt.Println("Address of return function from x variable: ", x)

}
