package main

func main() {

	//The func expression is assigments of funciton to variable

	x(10) // this will run function assigned to x

	y := func(num int) int {
		return num + 10
	}
	y(20) // this will run function assigned to y
}

var x = func(num int) int {
	return num + 10
}
