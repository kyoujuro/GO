package main

import "fmt"
//「math/rand」パッケージをインポート
import "math/rand"

func main() {
    
    for i := 1; i <= 10; i++ {
        // 0から99のランダムな整数を生成して
        fmt.Println(rand.Intn(100))
    }
}
