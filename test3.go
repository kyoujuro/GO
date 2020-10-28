package main

import(
    "fmt"
    "bufio"
    "log"
    "os"
)

func main(){
    fmt.Print("Enter a grade: ")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    log.Fatal(err)
    fmt.Println(input)
    
}
