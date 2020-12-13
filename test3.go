package main

import ("fmt")

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256}

func main(){
	x := complex(2.5, 3.1)
	fmt.Println(x)
	y := complex(5.1, 4.9)
	
	fmt.Println(y)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	//fmt.Println(cmplx.Abs(x))
	a := Axis{10, 20}
	fmt.Println(a.Y)

	member := [5]int{1, 2, 3, 4, 5}
	fmt.Println(member)
	fmt.Println(member[1])
	Factorial(pow)
}

type Axis struct {
	X int
	Y int
}

func Factorial(listA []int ){
	for i, v := range listA{
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
