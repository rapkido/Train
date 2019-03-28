package main

import (
	"fmt"
	"runtime"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go func() {
		ch <- doSomething(5)
		ch2 <- fmt.Sprint("CPU: ", runtime.NumCPU())
		ch3 <- fmt.Sprint("GOROUTINES ", runtime.NumGoroutine())

	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)
}
