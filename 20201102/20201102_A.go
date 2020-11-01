package main

import "fmt"

func Avg(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{1,3,5,7,9}
	avg := Avg(xs)
	fmt.Println(avg)
}