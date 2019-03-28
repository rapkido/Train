package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			fmt.Println(count)
			mu.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("Goroutines... \t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Final count is: ", count)
}
//race condition and without it exercise 3-4