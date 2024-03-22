package main

import "fmt"

func main() {
	fmt.Println("ARRAYS AND SLICES")

	//creating an array
	fmt.Println("Using arrays")
	// 1st posibility
	//array := []int{1, 2, 3, 4, 5}

	// 2nd posibility
	var array [5]int
	for i := 1; i <= len(array); i++ {
		array[i-1] = i
	}

	for index, val := range array {
		fmt.Printf("On index %v, stands value: %v, of type: %T\n", index, val, val)
	}

	fmt.Println("=================")

	//creating a slice literal
	// NOTE: slices are pointers to underlying array with len and capacity specified - just like dynamic arrays
	fmt.Println("Using Slices")
	// 1st posibility
	//slice := new([10]int)

	// 2nd posibility
	//slice := []int{1, 2, 3, 4, 5...}

	// 3nd posibility
	var slice []int = make([]int, 10)

	j := 0
	for j < len(slice) {
		slice[j] = j + 42

		j++
	}

	for index, val := range slice {
		fmt.Printf("On index %v, stands value: %v, of type: %T\n", index, val, val)
	}

	fmt.Println("=================")

	// slicing slices
	fmt.Println("Slicing slices")
	fmt.Printf("%v\n", slice[:5])
	fmt.Printf("%v\n", slice[6:])
	fmt.Printf("%v\n", slice[2:7])
	fmt.Printf("%v\n", slice[1:6])

	fmt.Println("=================")

	// appending to slices
	fmt.Println("Append to slice")
	slice = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice = append(slice, 52)

	slice = append(slice, 53, 54, 55)

	fmt.Printf("%#v\n", slice)

	slice2 := []int{56, 57, 58, 59, 60}

	slice = append(slice, slice2...)

	fmt.Printf("%#v\n", slice)

	fmt.Println("=================")

	// deleting to slices - use slicing and append to create shorter slice
	fmt.Println("Delete from slice")
	slice = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%#v\n", slice)

	// the "..." is for opening values in slice - like unwrapping
	slice = append(slice[:3], slice[6:]...)

	fmt.Printf("%#v\n", slice)

	fmt.Println("=================")

	// deleting to slices - use slicing and append to create shorter slice
	fmt.Println("Info about slice")

	fmt.Println("----below states len, capacity before and after----")

	//this initiates slice with 0 for all 50 indexes
	states := make([]string, 50)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(states[i], i)
	}

	fmt.Println("----below states2 len, capacity before and after----")

	// but this does not put any value in indexes
	states2 := make([]string, 0, 50)
	fmt.Println(len(states2))
	fmt.Println(cap(states2))

	states2 = append(states2, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(len(states2))
	fmt.Println(cap(states2))

	for i := 0; i < len(states2); i++ {
		fmt.Println(states2[i], i)
	}

	fmt.Println("=================")

	// two dimensional slices 
	fmt.Println("Two dimensional slices")
	
	bond := [][]string{{"James", "Bond", "Shaken, not stireed"}, {"Miss", "Moneypenny", "I'm 008"}}

	// this cloud also look like
	// jb := []string {"James", "Bond", "Shaken, not stireed"}
	// mm := []string {"Miss", "Moneypenny", "I'm 008"}
	// bond := [][]string{jb, mm}

	// loop over rows
	for i := range bond {
		fmt.Printf("%v\n", bond[i])
	}

	fmt.Println("=================")
}
