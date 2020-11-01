package main

import "fmt"

func main() {
	var min int
	x := []int{
		48, 97, 13, 59,
		35, 89, 10, 50,
		44, 29, 55, 9,
		90, 93, 11, 29,
	}
	for i, v := range x {
		if i == 0 || v < min {
			min = v
		}
	}
	fmt.Println(min)
}