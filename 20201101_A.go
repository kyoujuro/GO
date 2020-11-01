package main

import "fmt"

var y string = "hoge"
func main(){
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[0])
	fmt.Println("Hello " + "World")
	const x = "Good"
	fmt.Println("It is",x)
	fmt.Println(y)
	var input float64 = 23
	//fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
	fmt.Println(`1
	2
	3
	4
	5`)
	/*
	i := 1
	for i < 11 {
		fmt.Println("count is", i)
		i++
	}*/
	xs := []float64{10,40,33,4,23}
	total := 0.0
	for _, v := range xs{
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}

func f(){
	fmt.Println(y)
}