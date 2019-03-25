package main

import "fmt"

func main() {
	ss := []string{"a", "b", "c", "ç", "d", "e", "f", "g", "ğ", "h", "ı", "i", "j", "k", "l", "m", "n", "o", "ö", "p", "r", "s", "ş", "t", "u", "ü", "v", "y", "z"}
	for _, v := range ss {
		for _, va := range ss {
			gecici := fmt.Sprintf("%v%v", v, va)
			fmt.Println(gecici)
		}
	}
}
