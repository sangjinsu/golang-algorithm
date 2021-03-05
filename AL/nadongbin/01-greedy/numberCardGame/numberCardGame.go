package main

import "fmt"

func main() {
	cards := [][]int{
		{3, 1, 2},
		{4, 1, 4},
		{2, 2, 2},
	}
	fmt.Println(numberCardGame(cards))
}

func numberCardGame(cards [][]int) int {
	rowMin := 100001
	max := 0
	for _, rows := range cards {
		for _, card := range rows {
			if rowMin > card {
				rowMin = card
			}
		}
		if rowMin > max {
			max = rowMin
		}
		rowMin = 100001
	}
	return max
}
