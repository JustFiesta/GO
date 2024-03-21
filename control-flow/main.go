package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Printf("Setting stuff in init func\n")
}

func main() {
	x := rand.Intn(250) // exclusive at 250
	fmt.Printf("%v\n", x)

	//if statement
	if x >= 0 && x <= 100 {
		fmt.Printf("if - between 0 and 100\t x=%v\n", x)
	} else if x >= 101 && x <= 200 {
		fmt.Printf("if - between 101 and 200\t x=%v\n", x)
	} else {
		fmt.Printf("if - between 200 and 250\t x=%v\n", x)
	}

	//switch statement
	switch {
	case x >= 0 && x <= 100:
		fmt.Printf("switch - between 0 and 100\t x=%v\n", x)

	case x >= 101 && x <= 200:
		fmt.Printf("switch - between 101 and 200\t x=%v\n", x)

	case x <= 200:
		fmt.Printf("switch - between 200 and 250\t x=%v\n", x)

	default:
		fmt.Printf("switch - Adios")
	}
}
