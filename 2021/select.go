package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Min(a []int) (idx, n int) {	
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	
	return
}

func SelectionSort(array []int) []int {
	for i, _ := range array {
		idx, _ := Min(array[i:len(array)])
		array[i], array[i + idx] = array[i + idx], array[i]
	}
	
	return a
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	size := 50
	list := make([]int, size, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(1000)
	}
	
	fmt.Println(SelectionSort(list))
}
