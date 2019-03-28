package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Go Routines: ", runtime.NumGoroutine())
	//counter := 0
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	//sayac := 0
	wg.Add(gs)
	//var mu sync.Mutex We commented those old things because those weren't usefull
	start := time.Now()
	for i := 0; i < gs; i++ {
		//sayac++
		go func() {
			atomic.AddInt64(&counter, 1)
			//mu.Lock()
			//v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			//v++
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//counter = v
			//mu.Unlock()
			wg.Done()

		}()

		fmt.Println("Go Routines: ", runtime.NumGoroutine())

	}
	wg.Wait()
	endtime := time.Now()
	between := endtime.Nanosecond() - start.Nanosecond()
	fmt.Println("Time: ", between)
	//fmt.Println("STEPS: ", sayac) we get rid from sayac for examining code and it didn't changed
	fmt.Println("count: ", counter)
	fmt.Println("Hello, playground")
}

// 1993900,970700,995600 for mutex average 1320066,6666666666666666666666667 nanosec := 1,320066667 milisec
// 14960600,15921800,15925000 for just waitgroup and reace condition average 15602466,666666666666666666666667 nanosec:= 15,602466667 milisec
// 1996300,1995000, 1993500 for atomic level counting 1994933,3333333333333333333333333 nanosec := 1,994933333 milisec
