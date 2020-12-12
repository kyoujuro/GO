package main
import (
    "fmt"
    "strings"
)

func main(){
    broken := "G# r#cks!"
    replacer := strings.NewReplacer("#", "o")
    fixed := replacer.Replace(broken)
    fmt.Println(fixed)
    x := complex(2.5, 3.1)
	fmt.Println(x)
	y := complex(5.1, 4.9)
	fmt.Println(y)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
}

