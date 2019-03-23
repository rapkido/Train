package main

import (
	"fmt"
)

type square struct {
	L float64
	W float64
}
type circle struct {
	r float64
}

func (c circle) area() float64 {
	return 3.14159 * c.r * c.r
}
func (s square) area() float64 {
	return s.L * s.W
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	sq := square{
		L: 15,
		W: 15,
	}
	circ := circle{
		r: 12.345,
	}
	info(sq)
	info(circ)
	fmt.Println("Hello, playground")
}
