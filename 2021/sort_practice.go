package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		j := i - 1
		temp := s[i]
		for j > -1 && s[j] > temp {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = temp
	}
	return s
}

func shuffle(list []int) {
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i) // 0〜(i-1) の乱数発生
		list[i-1], list[j] = list[j], list[i-1]
	}
}

func MergeSort(s []int) []int {
	var result []int
	if len(s) < 2 {
		return s
	}

	mid := int(len(s) / 2)
	r := MergeSort(s[:mid])
	l := MergeSort(s[mid:])
	i, j := 0, 0

	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}

	result = append(result, r[i:]...)
	result = append(result, l[j:]...)

	return result
}

func BubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func Quicksort(s []int) []int {
	if len(s) == 1 || len(s) == 0 { //list size 0 or 1 is no sort
		return s
	} else {
		pivot := s[0] //make a pivot first in the list
		place := 0

		for j := 0; j < len(s)-1; j++ {
			if s[j+1] < pivot { // if it is smaller than the pivot
				s[j+1], s[place+1] = s[place+1], s[j+1]
				place++
			}
		}
		s[0], s[place] = s[place], s[0]

		first := Quicksort(s[:place])
		second := Quicksort(s[place+1:])
		first = append(first, s[place])

		first = append(first, second...)
		return first
	}
}

func main() {
	fmt.Println("OK")
	rand.Seed(time.Now().UnixNano())
	size := 50_000
	list := make([]int, size, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(1000)
	}
	shuffle(list)
	fmt.Println(list)
	//fmt.Println(InsertionSort(list))
	start := time.Now()
	MergeSort(list)
	end := time.Now()
	fmt.Printf("MergeSort %f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	BubbleSort(list)
	end = time.Now()
	fmt.Printf("BubbleSort %f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	Quicksort(list)
	end = time.Now()
	fmt.Printf("QuickSort %f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	InsertionSort(list)
	end = time.Now()
	fmt.Printf("InsertSort %f秒\n", (end.Sub(start)).Seconds())

}
