package main

import "fmt"

func main() {
	arr := []int{7, 8, 6, 1, 4, 3, 2, 0, 9, 5}

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	fmt.Println(arr)
}
