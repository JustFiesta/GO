package main

import (
	"fmt"
)

// puting 100 values to channel
func main() {
	// 1. make channel
	channel := make(chan int)

	// 2. launch goroutine for channel to run into
	go func() {
		for i := 0; i < 100; i++ {
			channel <- i // write nums to channel
		}
		close(channel) // close channel to prevent locking it further
	}() // NOTE: anonymus func should also prevent deadlocking channel

	// printing values from channel (no routine needed - it would read empty channel in concurrent fashion)
	for v := range channel {
		fmt.Println(v)
	}
}
