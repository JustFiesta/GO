package main

import "fmt"

func main() {

	// this channel is not in any running go routine (main doesn not count), which means I cant get things out of it
	// c := make(chan int)
	// c <- 42 // locks write on channel

	c := make(chan int)
	
	// now it should work - no deadlocks
	// this will lock channel until it writes to it
	go func(){
		c <- 42
	}()
	
	// this lock channel until it pulls somehing from channel
	fmt.Println(<-c)

	// to sum it up:
	// channels must meet after one does something and unlocks channel
	// just like in relay-race


	//2n solution to deadlock is using buffered channels - ones with predefined capacity
	c2 := make(chan int, 1) // channel with buffer capacity 1

	c2 <- 44 // pass int to channel

	fmt.Println(<-c2) // output from channel - not ideal bcs it need predefined channel capacity
}
