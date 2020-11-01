package main

import "fmt"

type Person struct {
    first_name string
    age int 
}

func newPerson(first_name string, age int) *Person{
    person := new(Person)
    person.first_name = first_name
    person.age = age
    return person
}




func main(){
    var tom *Person = newPerson("Tom", 22)
    fmt.Println(tom.first_name, tom.age)
}

    
