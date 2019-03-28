package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	onething := "Bu birincisi!!"
	secondthing := "Bu ikincisi!!"
	go func(thing string) {
		fmt.Println(thing)
		wg.Done()
	}(onething)
	go func(thing string) {
		fmt.Println(thing)
		wg.Done()
	}(secondthing)
	wg.Wait()
}
