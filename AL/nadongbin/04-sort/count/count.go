package main

import "fmt"

func main() {
	arr := []int{7, 8, 6, 1, 4, 3, 2, 0, 9, 5, 8, 6, 2, 1, 3, 4, 5, 6}

	countList := make([]int, 10)

	for _, value := range arr {
		countList[value]++
	}

	fmt.Println(countList)

	sortedArr := []int{}
	for i, count := range countList {
		for j := 0; j < count; j++ {
			sortedArr = append(sortedArr, i)
		}
	}

	fmt.Println(sortedArr)
}
