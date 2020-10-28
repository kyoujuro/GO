package main

import "fmt"
import "time"
import "math/rand"
func main() {
    a := []int{1, 2, 3}
    fmt.Println(choice(a))
    fmt.Println(choice(a))
    k := make_array(50)
    fmt.Println(k)

    fmt.Println(bubbleSort(k))
    fmt.Println(bubbleSort2(k))
}

func choice(s []int) int {
    rand.Seed(time.Now().UnixNano())
    i := rand.Intn(len(s))
    return s[i]
}

func bubbleSort(array []int) []int{
    for i := 0; i < len(array)-1; i++{
        for j := i; j < len(array); j++{
            if array[i] < array[j]{
                tmp := array[i]
                array[i] = array[j]
                array[j] = tmp 
            }
        }
    }
    return array
}


func bubbleSort2(array []int) []int{
    for i := 0; i < len(array)-1; i++{
        for j := i; j < len(array); j++{
            if array[i] > array[j]{
                tmp := array[i]
                array[i] = array[j]
                array[j] = tmp 
            }
        }
    }
    return array
}
func make_array(num int) []int{
    var array []int
    for i := 0; i < num; i++ {
        rand.Seed(time.Now().UnixNano())
        array = append(array, rand.Intn(1000))
    }
    return array
}