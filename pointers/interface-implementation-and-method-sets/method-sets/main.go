package main

import "fmt"

type person struct {
	first string
}

// method for type pointer to person
func (p *person) speak() {
	fmt.Println("My name is ", p.first)
}

// every type that with speak method is a human
type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}


func main() {
	
	p := person{
		first: "Fifonż",
	}
	
	// checking if can pass a type *person (pointer to person) to saySomething (i should according to go method sets - https://go.dev/ref/spec#Method_sets)
	
	saySomething(&p) // - no errors


	// checking if can pass a type person (value) to saySomething (i should according to go method sets - https://go.dev/ref/spec#Method_sets)
	
	saySomething(p) // - error: speak uses pointer reciver (func (t *T))

	// buuut this works
	p.speak()

	// Conclusion:
	// This situation is becouse interface does not recognise speak with value reciver - one does not exist => person is not human
	// but as a method for person, speak do not care about it - more in link above
}