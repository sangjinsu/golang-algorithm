package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 1, 2, 2, 2, 2, 3}

	left := sort.SearchInts(arr, 2)
	right := sort.Search(len(arr), func(i int) bool { return arr[i] > 2 })

	fmt.Println(right - left)
}
