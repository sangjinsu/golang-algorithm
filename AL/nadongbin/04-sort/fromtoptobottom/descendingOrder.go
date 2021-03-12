package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, 3)
	var value int

	for i := 0; i < n; i++ {
		fmt.Scanln(&value)
		arr[i] = value
	}

	// for i := 1; i < n; i++ {
	// 	for j := i; j > 0; j-- {
	// 		if arr[j-1] < arr[j] {
	// 			arr[j-1], arr[j] = arr[j], arr[j-1]
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
}
