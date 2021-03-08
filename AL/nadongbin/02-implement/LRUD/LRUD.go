package main

import (
	"fmt"
	"strings"
)

func showDestination(n int, moves string) string {
	movesArr := strings.Split(moves, " ")

	point := []int{1, 1}

	// U D L R
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for _, move := range movesArr {
		x, y := point[0], point[1]

		switch move {
		case "L":
			x += dx[2]
			y += dy[2]
		case "R":
			x += dx[3]
			y += dy[3]
		case "U":
			x += dx[0]
			y += dy[0]
		case "D":
			x += dx[1]
			y += dy[1]
		}

		if x < 1 || x > n || y < 1 || y > n {
			continue
		}

		point[0], point[1] = x, y
	}
	return fmt.Sprint(point[0]) + " " + fmt.Sprint(point[1])
}

func main() {
	fmt.Println(showDestination(5, "R R R U D D"))
}
