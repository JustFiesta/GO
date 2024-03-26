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

// change q to recive only
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	// this locks channel c - deadlock (it still waits for icoming data)
	// for i := 0; i < 100; i++ {
	// 	c <- i
	// } // need to create a anonymus func and close channel after we are done

	go func() { // send channel to its routine
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1 // put 1 into q chan
		close(c) // closing channel after recibving data
	}()

	return c
}

func receive(c, q <-chan int) {
	var v int

	// forever check if channels have some value
	for {
		select {
			case v = <-c: // pull value from c channel
				fmt.Println(v)
			case <-q: // in case that we wanna quit (channel quit is empty - it defines that we need to close select)
				fmt.Println("quit")
				return
			}
	}
}
