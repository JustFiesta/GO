package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func ()  { // create goroutine for channel to start working
		c <- 44
	}()

	v, ok := <-c // if we get value pass value
	fmt.Println(v, ok)

	close(c)
	
	v, ok = <-c // this shloud have 0 value (bcs channel is closed) and ok should be false
	fmt.Println(v, ok) // this shows how comma ok idiom works
}
