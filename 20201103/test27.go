package main


import (
    "fmt"
    "sort"
)


func main(){
    strs := []string{"n", "b", "a"}
    sort.Strings(strs)
    fmt.Println("Strings", strs)
     
    ints := []int{1, 20, 7}
    t := sort.IntsAreSorted(ints)
    fmt.Println("Sorted:", t)
    sort.Ints(ints)
    fmt.Println("Ints", ints)
  
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted:", s)
}
