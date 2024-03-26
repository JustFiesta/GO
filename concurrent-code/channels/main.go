// Put and pull values into channel
package main

import "fmt"

// create 10 goroutines - each should put 10 values into channel
func main()  {

	channel := make(chan int)

	for i := 0; i < 10;i++ {
		go func () {
			for j := 0; j < 10; j++{
				channel <- j
			}
			// close(channel) // if closed here - other routines might not be able to put data
		}()
	}

	// num of values in channel is known - so no need to close and reopen channel
	// for v := range channel {
	// 	fmt.Println(v)
	// }
	for k := 0; k < 100;k++ {
		fmt.Println(<-channel)
	}
}