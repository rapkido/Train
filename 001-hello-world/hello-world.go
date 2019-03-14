package main

import "fmt"

func main() {
	fmt.Println("Hello TODD thank you for that lesson i will be real prof programmer through you")
	foo()
	fmt.Println("After foo it's something like this you know")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()

}
func foo() {
	fmt.Println("i am in foo")
}
func bar(){
	fmt.Println("It's end of code")
}

// control flow
// 1 sequence
// 2 loop;itterative
// 3 conditional
