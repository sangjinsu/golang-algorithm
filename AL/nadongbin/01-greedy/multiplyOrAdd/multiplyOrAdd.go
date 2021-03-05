package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiplyOrAdd(numStr string) int {
	numChars := strings.Split(numStr, "")
	nums := make([]int, len(numChars))
	for i, nc := range numChars {
		nums[i], _ = strconv.Atoi(nc)
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		next := nums[i]

		if result < 2 || next < 2 {
			result += next
		} else {
			result *= next
		}

		fmt.Println(result)
	}

	return result
}

func main() {
	multiplyOrAdd("12984")
}
