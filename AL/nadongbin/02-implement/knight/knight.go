package main

import "fmt"

func countMoves(location string) int {
	x, y := int(location[1]-byte('0')), int(location[0]-byte('a'))+1

	dx := []int{1, 2, 2, 1, -1, -2, -2, -1}
	dy := []int{-2, -1, 1, 2, 2, 1, -1, -2}

	count := 0

	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 1 || nx > 8 || ny < 1 || ny > 8 {
			continue
		}
		count++
	}

	return count
}

func main() {
	fmt.Println(countMoves("b3"))
}
