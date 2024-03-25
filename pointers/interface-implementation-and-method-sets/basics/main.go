package main

import "fmt"


// this type and its methods are immutable in this excercise
type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walk.")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

// instead of using pointer semantics choosed to do value semantics - according to its heuristics its better in this excercise
type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)
}

// mutate value by function call instead of pointers
func (d dog) changeName(s string) dog {
	d.first = s
	return d
}