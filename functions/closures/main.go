package main

import (
	"fmt"
	"math"
)

func powerinator(a float64) func() float64{
	var c float64

	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

func main() {
	// Closure function is a function that encloses other function in the body
	// after assigment to variable, every call of this function the c will be incremented

	x := powerinator(2)

	// use parenthesis here to run enclosed function with increments c
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}
