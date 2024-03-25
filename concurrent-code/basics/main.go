package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func main runs its own goroutine
func main() {

	fmt.Println("=========Goroutines in background===========")

	// in addition to that I'm creating two new ones
	go hello()
	go world()

	// and to check how they are doing in runtime I use runtime package

	runtime.NumGoroutine()
	runtime.NumCPU()

	fmt.Println("=========Goroutines in foreground===========")

	// this will not output anything as they run in background (same as command& in shell)
	// to check really how it is going just use loop

	fmt.Println("CPU cores used: ", runtime.NumCPU())
	fmt.Println("Go routines - start: ", runtime.NumGoroutine()) // there should be three as - two from above + main func
	for i := 0; i < 5; i++ {

		go hello()
		go world()

	}
	fmt.Println("Go routines - stop: ", runtime.NumGoroutine())

	// to make goroutines finish before program ends - need to somehow wait for them to end; Sync package (waitgroups/mutex/atom)

	fmt.Println("========WaitGroups==========")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		wg.Add(2) // wait for two routines to finish

		go func(){
			hello()
			wg.Done() // decrement .Add by one
		}()
		go func(){
			world()
			wg.Done()// decrement .Add by one - all routines are done
		}()
	}
	wg.Wait() // wait for all routines to end
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world!")
}
