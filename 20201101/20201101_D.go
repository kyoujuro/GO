package main

import (
	"crypto/sha1"
	"fmt"
)

func main(){
	h := sha1.New()
	h.Write([]byte("hoge"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}