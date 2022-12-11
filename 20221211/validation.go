package main

import "fmt"
import "regexp"

func main() {
	timestamp_validation("2022/12/12")
}

func timestamp_validation(timestamp string) {
	re := regexp.MustCompile(`(\d{4})/(\d{2})/(\d{2})`)
	result := re.MatchString(timestamp)
	return result
}
