// Print OS and ARCH (Architecture) info
package main

import (
	"fmt"
	"runtime"
)


// this program is made for demonstrating purposes of go: run, build, install commands
func main() {
	//for env info check go spec/go env/runtime spec
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)
}