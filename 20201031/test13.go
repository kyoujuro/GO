package main

import "fmt"


type Person struct{
    first_name string
    age int
}


func main(){
    var mike Person
    var bob Person
    mike.first_name = "Mike"
    mike.age = 10
    bob := Person{"bob", 20}
    fmt.Println(mike.first_name, mike.age,bob.age)
}

