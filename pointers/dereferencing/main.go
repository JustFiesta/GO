package main

import "fmt"

var (
	a, b, c *string //this points to an address where stuff will be stored
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p // example here - we assing the address of p to pointer
	b = &q
	c = &r
	d = &n
}

func main() {
	// pointers store address, to get value behind it we need dereferencing
	fmt.Printf("The memory address of a is: %v. It's type is: %T. And the value they hold is: %v\n", a, a, *a) // without an %v this prints every bit of information about type, where with it it prints just the value that pointer holds
	fmt.Printf("The memory address of b is: %v. It's type is: %T. And the value they hold is: %v\n", b, b, *b) // without an %v this prints every bit of information about type, where with it it prints just the value that pointer holds
	fmt.Printf("The memory address of c is: %v. It's type is: %T. And the value they hold is: %v\n", c, c, *c) // without an %v this prints every bit of information about type, where with it it prints just the value that pointer holds
	fmt.Printf("The memory address of d is: %v. It's type is: %T. And the value they hold is: %v\n", d, d, *d) // without an %v this prints every bit of information about type, where with it it prints just the value that pointer holds
}
