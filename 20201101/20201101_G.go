package main

import "fmt"

func zero(xPtr *int){
	*xPtr = 6
}

func change(x int){
	x = 4
}

func main(){
	x := 5
	zero(&x)
	fmt.Println(x)
	change(x)
	fmt.Println(x)
	y := new(int)
	change(*y)

	fmt.Println(y)
}