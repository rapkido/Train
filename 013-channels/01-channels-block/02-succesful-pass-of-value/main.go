package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func (){
		fmt.Println("Sleeping...")
		time.Sleep(2*time.Second)
		c <-42
			}()

	fmt.Println(<-c)
}

