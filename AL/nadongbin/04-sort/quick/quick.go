package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := start
	left := start + 1
	right := end

	for left <= right {
		for left <= end && arr[left] <= arr[pivot] {
			left++
		}
		for right >= start && arr[right] >= arr[pivot] {
			right--
		}
		if left > right {
			arr[right], arr[pivot] = arr[pivot], arr[right]
		} else {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	quickSort(arr, start, right-1)
	quickSort(arr, right+1, end)
}

func main() {
	arr := []int{6, 5, 8, 1, 3, 2, 9, 4, 7, 0}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
