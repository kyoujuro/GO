package main

import "fmt"

func main([]int array)[] int{
  for i := 0; i < len(array); i++{
    for j := i; j < i; j++{
      if array[i] < array[j]{
        array[i], array[j] = array[j], array[i]
      }
    }
  }
  return array
}
