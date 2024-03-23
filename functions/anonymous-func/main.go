package main

import "fmt"

func main() {

	//func(params)(return type){ code }(arguments)
	func(num int) int {
		fmt.Println("Appending 10 to num")
		return num + 10
	}(20)
}