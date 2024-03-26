package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	// for i := 0; i < 100; i++ {
	// 	c <- i
	// } // ... if closed here it will return locked at reciving channel

	go func() {
		for i := 0; i < 100; i++ {
			c <- i 
		}
		close(c) // close it after goroutine is don sending
	}()

	return c
}

// function to receve values from certain channel, using for range loop
func receive(c, q <-chan int) {
	for v := range c {
		fmt.Println(v) // c is still open here. need to close it...
	}
	// ... second option is to make an anonymus func and close after it finishes
}
