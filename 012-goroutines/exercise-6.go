package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Your sperating sistem is: ",runtime.GOOS)
	fmt.Println("Your operating system is: ",runtime.GOARCH)
}
