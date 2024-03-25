package main

import (
	"fmt"
	"runtime"
	"sync"
)

// eg. in concurrent code two functions/modules/stuff, might read data from same place in memory

var site_entries int

func web_bot() {

	v := site_entries

	runtime.Gosched()

	v++
	site_entries = v
}

func user_entry() {

	v := site_entries

	runtime.Gosched()

	v++
	site_entries = v
}

func main() {

	fmt.Println("Initial site entries value: ", site_entries)

	var waitgroup sync.WaitGroup

	// this shloud create 10 goroutines, all want to increment site_entires - so it shloud be 10 by end
	for i := 0; i < 5; i++ {
		waitgroup.Add(2) // waiting for two runtimes

		// creating callback funcs to run goroutines
		func () {
			go web_bot() // run goruntime
			waitgroup.Done() // decrement waiting list from .Add
		}()

		func () {
			go user_entry() // run goruntime
			waitgroup.Done() // decrement waiting list from .Add
		}()
	}

	waitgroup.Wait()
	
	fmt.Println("Final site entries value: ", site_entries)

	// run this and see that its not consistant

	// this is called data race condition - check it with go run -race main.go
}