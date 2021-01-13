package main

import ("fmt"
        "math/rand"
        "time"
)

func BubbleSort(array []int) []int{
  for i := 0; i < len(array) ; i++{
    for j := i; j <len(array) ; j++{
      if array[i] > array[j]{
        array[i], array[j] = array[j], array[i]
      }
    }
  }
  return array
}

func main(){
  rand.Seed(time.Now().UnixNano())
	size := 100
	array := make([]int, size, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(50)
	}
  fmt.Println(BubbleSort(array))
}
