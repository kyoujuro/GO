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

type  SerialPort_Config  struct{
    Serial_path string,
    Rate int,
    Bits int,
    Stop int,
    Prity string,
}
