package main

import "fmt"

func do(i interface{}) {
    switch variable := i.(type){
        case int:
            fmt.Println(variable)
        case string:
            fmt.Println(variable)
        default:
            fmt.Println("Default")
        }
}


func main(){
    do(23)
    do("hoge")
    do(true)
}
 
