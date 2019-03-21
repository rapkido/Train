package main

import (
	"fmt"
)

func main() {
	s := struct {
		m     map[int]string
		datas []int
	}{
		m: map[int]string{
			94: "besto",
			57: "baba",
			63: "anne"},
		datas: []int{57, 63, 94},
	}
	fmt.Println(s.m[s.datas[2]])

}
