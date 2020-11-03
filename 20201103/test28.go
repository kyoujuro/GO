package main

import (
	"fmt"
	"encoding/json"
//	"os"
)
type responce1 struct {
	Page int
	Fruites []string
}
type responce2 struct {
	Page int `json:"page"`
	Fruites []string `json:"fruites`
}

func main(){

	bolB, _ := json.Marshal(1)
	fmt.Println(string(bolB))
	
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	slcD := []string{"apple","banana","peach"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
}
