package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func mergeSort(m []int) []int {
	if len(m) == 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func mergeSortMulti(m []int) []int {
	if len(m) == 1 {
		return m
	}

	var wg sync.WaitGroup
	wg.Add(2)

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	arr := make(chan []int, 2)
	go func(left []int, arr chan []int) {
		defer wg.Done()
		arr <- mergeSortMulti(left)
	}(left, arr)

	go func(right []int, arr chan []int) {
		defer wg.Done()
		arr <- mergeSortMulti(right)
	}(right, arr)
	wg.Wait()

	return merge(<-arr, <-arr)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	leftIdx, rightIdx := 0, 0
	idx := 0
	for len(left) > leftIdx || len(right) > rightIdx {
		if len(left) > leftIdx && len(right) > rightIdx {
			if left[leftIdx] <= right[rightIdx] {
				result[idx] = left[leftIdx]
				leftIdx += 1
			} else {
				result[idx] = right[rightIdx]
				rightIdx += 1
			}
		} else if leftIdx < len(left) {
			result[idx] = left[leftIdx]
			leftIdx += 1
		} else if rightIdx < len(right) {
			result[idx] = right[rightIdx]
			rightIdx += 1
		}
		idx += 1
	}
	return result
}

func main() {
	start := time.Now()
	fmt.Println(mergeSort([]int{47, 84, 37, 29, 4, 45, 93, 20, 98, 49, 67, 73, 100, 71, 52, 78, 85, 96, 39, 25, 27, 31, 24, 32, 26, 86, 40, 44, 64, 75, 70, 30, 14, 77, 28, 5, 19, 97, 41, 22, 11, 60, 9, 1, 7, 59, 72, 61, 82, 65}))
	elapsed := time.Since(start)
	log.Println(elapsed.Microseconds())

	start = time.Now()
	fmt.Println(mergeSortMulti([]int{47, 84, 37, 29, 4, 45, 93, 20, 98, 49, 67, 73, 100, 71, 52, 78, 85, 96, 39, 25, 27, 31, 24, 32, 26, 86, 40, 44, 64, 75, 70, 30, 14, 77, 28, 5, 19, 97, 41, 22, 11, 60, 9, 1, 7, 59, 72, 61, 82, 65}))
	elapsed = time.Since(start)
	log.Println(elapsed.Microseconds())
}
