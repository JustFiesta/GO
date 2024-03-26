package main

import (
	"fmt"
)

func main() {
	// send only channel (read from left: channel (chan) for sending to it (<-) ints (int))
	// cs := make(chan<- int) - this couses issues with reading its values (kind of pipeline needed for reading)
	cs := make(chan int) // instead changed its type to unidirectional channel

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	// receve only channel (read from left: sending to cr (<-) channel data (chan) ints (int))
	// cr := make(<-chan int)
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
