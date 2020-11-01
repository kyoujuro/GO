package main

import "fmt"

type Person struct {
    first_name  string
    age int
}

func(p Person) intro(greetings string) string{
    return greetings + " I am " + p.first_name
}

func main(){
    bob := Person{"Bob", 30}
    fmt.Println(bob.intro("Hello"))
}
    
