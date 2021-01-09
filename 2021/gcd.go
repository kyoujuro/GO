package main

import "fmt"

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func main() {
    fmt.Println(gcd(29024, 7964)) // 4
}
