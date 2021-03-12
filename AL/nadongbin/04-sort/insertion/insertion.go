package main

import "fmt"

func main() {
	arr := []int{7, 8, 6, 1, 4, 3, 2, 0, 9, 5}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

	fmt.Println(arr)
}
