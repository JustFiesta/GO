package main

import "fmt"

var X int = 42
const Y string = "before/outside main scope"

func main()  {
	z := "inside main scope"

	fmt.Printf("Hello! %v. %v", Y, z)
}