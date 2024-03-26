package main

import (
	"fmt"
)

type customErr struct {
	info string
}

// creating method Error so person can implement Error type
func (ce customErr) Error() string {

	return "This is my custom error: " + ce.info
}

func foo(err error) {
	fmt.Println("This is your error: ", err)
}

func main() {
	ce := customErr{
		info: "custom error stuff",
	}

	// bcs customError implements Error interface it can be passed to foo!
	foo(ce)
}