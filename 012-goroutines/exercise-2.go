package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (s *person) speak() {
	fmt.Println(s.First)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

//you CAN pass a value of type *person into saySomething
//you CANNOT pass a value of type person into saySomething
func main() {
	p1 := person{
		First: "Bestami",
		Last:  "Coskuner",
		Age:   25,
	}
	saySomething(&p1)
	(&p1).speak()
}
