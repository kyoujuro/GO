package main

import ("fmt"
		"math"
)

func main(){
	fmt.Println(Golden(10))	
}

func Golden(n int) float64 {
    return math.Round((math.Pow((1+math.Sqrt(5))/2, float64(n)) - 
		       math.Pow((1-math.Sqrt(5))/2, float64(n))) / math.Sqrt(5))
}
