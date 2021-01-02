package main

import (
    "fmt"
    "math"
)


func max(lhs, rhs int) int {
    return int(math.Max(float64(lhs), float64(rhs)))
}

func main() {
    var (
        N, M int
    )
    // 入力
    fmt.Scanf("%d %d", &N, &M)
    values, weights := make([]int, N), make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d %d", &weights[i], &values[i])
    }
  
    dp := make([][]int, N+1)
    for i := 0; i < N+1; i++ {
        dp[i] = make([]int, M+1)
    }
   
    for i := 1; i <= N; i++ {
        for j := int(0); j <= M; j++ {
            dp[i][j] = dp[i-1][j]
            if j >= weights[i-1] {
                dp[i][j] = max(dp[i][j], dp[i-1][j-weights[i-1]]+values[i-1])
            }
        }
    }
    // 出力
    fmt.Println(dp[N][M])
}
