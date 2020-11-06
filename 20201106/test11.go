package main

import "fmt"

func main(){
   fmt.Println("YES")
   defer fmt.Println("NO")
   fmt.Println("go")
}
