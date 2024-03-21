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
		fmt.Printf("switch - Adios\n")
	}

	//for loop - there are 3 types of them (find more on go spec/effective go)

	// C-like for
	// for i := 0; i < count; i++ {
	//		stuff
	// }

	// C while-like for
	// for i < 10 {
	// 	stuff
	// 	i++
	// }

	// looping over array/slices/dicts/etc
	// for key, value := range oldMap {
	// 	newMap[key] = value
	// }

	// Above one has some variations
	// If you only need the first item in the range (the key or index), drop the second:
	// for key := range m {
	// 	if key.expired() {
	// 		delete(m, key)
	// 	}
	// }
	// If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:

	// sum := 0
	// for _, value := range array {
	// 	sum += value
	// }

	fmt.Printf("=====================\n")

	var i int
	for i < 100 {
		fmt.Printf("for - Num of iterations: %v\n", i+1)

		x := rand.Intn(250)

		if x >= 0 && x <= 100 {
			fmt.Printf("between 0 and 100\t x=%v\n", x)
		} else if x >= 101 && x <= 200 {
			fmt.Printf("between 101 and 200\t x=%v\n", x)
		} else {
			fmt.Printf("between 200 and 250\t x=%v\n", x)
		}

		i++
	}

	fmt.Printf("=====================\n")

	i = 0
	for i < 42 {
		fmt.Printf("for switch - Num of iterations: %v\n", i+1)

		x = rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("Got 0\n")
		case 1:
			fmt.Printf("Got 1\n")
		case 2:
			fmt.Printf("Got 2\n")
		case 3:
			fmt.Printf("Got 3\n")
		case 4:
			fmt.Printf("Got 4\n")
		}

		i++
	} 

	fmt.Printf("=====================\n")

	i = 0
	// infinite loop
	for {
		x = rand.Intn(5)
		fmt.Printf("infinite - Num of iterations: %v\n", i+1)
		if x == 4 {
			break
		}
		i++
	}

	fmt.Printf("=====================\n")

	// print all odd numbers
	fmt.Printf("Odd nums to 100\n")
	for i := 1; i <= 100; i++ {
		if ! (i % 2 == 1) {
			continue
		}
		fmt.Printf("odd: %v\n", i)
	}

	fmt.Printf("=====================\n")

	// nested loops
	fmt.Printf("Nested loops\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("Outside loop - iteration: %v\n", i + 1)
		for j := 0; j < 5; j++ {
			fmt.Printf("Nested loop - iteration: %v\n", j + 1)
		}
	}

	fmt.Printf("=====================\n")

	// iterate over slice - using range
	fmt.Printf("Iterate over slice\n")
	xi := []int{42, 42, 44, 45, 46, 47}
	for index, value := range xi {
		fmt.Printf("On index: %v is value of: %v\n", index, value)
	}

	fmt.Printf("=====================\n")

	// iterate over map - using range
	fmt.Printf("Iterate over map\n")
	yi := map[string]int{
		"Fifonż": 12,
		"Alfons": 53,
	}
	for key, value := range yi {
		fmt.Printf("On key: %v is number of: %v\n", key, value)
	}

	fmt.Printf("=====================\n")

	// if statement; ok; statement
	age := yi["Fifonż"]
	fmt.Printf("The age of Fifonż is: %v\n", age)

	fmt.Printf("Use if statement; ok; statement\n")
	age2 := yi["Q"]
	fmt.Printf("The age of Q is: %v\n", age2)

	if _, ok := yi["Q"]; !ok {
		fmt.Printf("There is no Q in your map\n")
	}

	fmt.Printf("=====================\n")

	// if statement; statement
	fmt.Printf("Use if statement; statement\n")
	i = 0
	for i < 5{
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("x is %v", x)
		}
		i++
	}

	fmt.Printf("=====================\n")

	// checking binary operators
	fmt.Println(true && true) 
	fmt.Println(true && false) 
	fmt.Println(true || true) 
	fmt.Println(true || false) 
	fmt.Println(!true) 
}
