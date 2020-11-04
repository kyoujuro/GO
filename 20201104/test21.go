package main

import ("fmt"
        "time"
)


func main(){
    fmt.Println("start")
    process(2,"A")
    process(2,"B")
    fmt.Println("finish")
}


func process(num int, str string){
    for i := 0; i <= num; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println(i, str)
    }
}

