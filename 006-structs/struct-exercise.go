package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	volvo := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: false,
	}
	mondeo := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println("volvo")
	fmt.Println(volvo.doors, volvo.color, volvo.fourWheel)
	fmt.Println("mondeo")
	fmt.Println(mondeo.doors, mondeo.color, mondeo.luxury)
}
