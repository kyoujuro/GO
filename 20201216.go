package main
import "fmt"

func selection_sort(ar []int) []int {
    for i := 0; i < len(ar) - 1; i++ {
        min := i
        for j := i + 1; j < len(ar); j++ {
            if ar[j] < ar[min] { min = j }
            ar[min], ar[i] = ar[i], ar[min]
        }
    }
    return ar
}

func main() {
    ar := selection_sort([]int{})
    fmt.Println(ar)
}
