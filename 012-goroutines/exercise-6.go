package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Your operating sistem is: ", runtime.GOOS)
	fmt.Println("Your system architechture is: ", runtime.GOARCH)
}
