package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// fixing this code using atomic - low level locking for certain type pointer

// incrementor for atomic methods (with pointers)
var site_entries int64 // need correct format - this will be pointer to some memory location for atomic

func main() {

	fmt.Println("Initial site entries value: ", site_entries)

	var waitgroup sync.WaitGroup // at first - synchronize goroutines

	waitgroup.Add(5) // set count for each one

	// this shloud create 5 goroutines, all want to increment site_entires - so it shloud be 5 by end
	for i := 0; i < 5; i++ {

		// creating callback funcs to run goroutines
		go func() {

			//use of atomic to increment number
			atomic.AddInt64(&site_entries, 1) // bcs this increments certain memory location three below lines are useless
			
			// v := site_entries
			// v++
			// site_entries = v

			waitgroup.Done() // mark routine as done
		}()

	}

	waitgroup.Wait()

	fmt.Println("Final site entries value: ", site_entries)
}
