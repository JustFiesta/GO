package main

import "fmt"

type person struct {
	first string
	age int
}

// methods differ in decalration - func (variable Type) name() {}
func (p person) speak() {
	fmt.Printf("My name is %v and my age is %v", p.first, p.age)
	//fmt.Println("My name is ", p.first, " and my age is ",  p.age)
} 

func main() {


	p := person{
		first: "Fifonż",
		age: 26,
	}

	p.speak()
}
