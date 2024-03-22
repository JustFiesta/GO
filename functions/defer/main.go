package main

import "fmt"

func main() {

	// defer waits for other functions to finish and runs lastly called function

	defer fmt.Println("Yoooooo!")

	fmt.Println("Hiii!")

	defer fmt.Println("World!")

	fmt.Println("Adios!")

	// it runs funcions as LIFO - last in first out
	// output:
// 	Hiii!
//	Adios!
// 	World!
// 	Yoooooo!

}
