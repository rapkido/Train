package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count,1)
			r:=atomic.LoadInt64(&count)
			fmt.Println(r)
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Final count is: ", count)
}
