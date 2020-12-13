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
    var testA map[string]string = map[string]string{"Tokyo":"Tokyo", "Aichi":"Nagoya"}
    fmt.Println(testA["Aichi"])
    testB := make(map[string]string)
    fmt.Println(testB)
    testC := make(map[string]string, 20)
    fmt.Println(testC)
    a := Axis{10, 20}
    fmt.Println(a.Y)
    member := [5]int{1, 2, 3, 4, 5}
    fmt.Println(member)
    fmt.Println(member[1])
}

type  SerialPort_Config  struct{
    Serial_path string,
    Rate int,
    Bits int,
    Stop int,
    Prity string,
}
