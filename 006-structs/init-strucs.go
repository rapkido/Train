package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "Bestami",
		last:  "Coskuner",
	}
	p2 := person{
		first: "Aram",
		last:  "Alp",
	}
	p3 := secretAgent{
		person: person{first: "James", last: "Bond"},
		ltk:    true}
	fmt.Println(p1, p2, p3)

}
