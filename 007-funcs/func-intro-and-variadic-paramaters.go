package main

import "fmt"

func sum(n ...float64) (sum, l float64) {
	fmt.Println("The slice is ", n)
	sum = 0
	for i, v := range n {
		sum += v
		fmt.Println("We are adding the element which stored in slices with", i, "index", "and now sum is equal to ", sum)
	}
	return sum, float64(len(n))
}
func main() {
	x, l := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The sum of slices is ", x)
	fmt.Println("Average of this slicesi is", x/l)
}
