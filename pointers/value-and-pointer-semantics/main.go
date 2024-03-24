package main

import "fmt"

type person struct {
	first  string
	second string
}

// no need to return - the person behind pointer is changed
func changeNameP(p *person, s string){
	p.first = s
}

// need to return changed value - GO uses copys of values not the same references!!!
func changeNameV(p person, s string) person {
	p.first = s
	return p
}

func main() {
	p := person{
		first: "Fizonż",
		second: "Nowakowski",
	}

	p2 := person{
		first: "Gerwazy",
		second: "Nowaki",
	}

	p = changeNameV(p, "Alfons")

	changeNameP(&p, "Alfred")

	fmt.Println(p)
	fmt.Println(p2)
}
