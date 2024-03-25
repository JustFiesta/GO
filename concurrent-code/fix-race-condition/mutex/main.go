package main

import (
	"fmt"
	"sync"
)

// fixing this code using mutex - high level locking for certain parts of program

var site_entries int

// mutual exclusion lock 
var mut sync.Mutex

func main() {

	fmt.Println("Initial site entries value: ", site_entries)

	var waitgroup sync.WaitGroup // at first - synchronize goroutines

	waitgroup.Add(5) // set count for each one

	// this shloud create 5 goroutines, all want to increment site_entires - so it shloud be 5 by end
	for i := 0; i < 5; i++ {

		// creating callback funcs to run goroutines
		go func() {
			mut.Lock() // locks chunk of code

			v := site_entries
			v++
			site_entries = v
		
			mut.Unlock() // unlocks chunk of code

			waitgroup.Done() // mark routine as done
		}()

	}

	waitgroup.Wait()

	fmt.Println("Final site entries value: ", site_entries)
}
