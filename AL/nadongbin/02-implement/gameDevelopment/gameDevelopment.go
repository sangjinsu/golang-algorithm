package main

import "fmt"

func gameDevelopment(mapSize [2]int, gamer [3]int, m [][]int) int {
	var mapHistory [][]int
	zero := []int{}

	for i := 0; i < mapSize[1]; i++ {
		zero = append(zero, 0)
	}

	for i := 0; i < mapSize[0]; i++ {
		zeroArr := make([]int, len(zero))
		copy(zeroArr, zero)
		mapHistory = append(mapHistory, zeroArr)
	}

	x, y, direction := gamer[0], gamer[1], gamer[2]

	mapHistory[x][y] = 1

	// 북 동 남 서
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	count := 1
	turnTime := 0
	for {
		direction = turnLeft(direction)

		nx, ny := x+dx[direction], y+dy[direction]

		if mapHistory[nx][ny] == 0 && m[nx][ny] == 0 {
			mapHistory[nx][ny] = 1
			x = nx
			y = ny
			count++
			turnTime = 0
			continue
		} else {
			turnTime++
		}

		if turnTime == 4 {
			nx = x - dx[direction]
			ny = y - dy[direction]

			if m[nx][ny] == 0 {
				x = nx
				y = ny
			} else {
				break
			}

			turnTime = 0
		}
	}

	return count
}

func turnLeft(direction int) int {
	direction--
	if direction == -1 {
		direction = 3
	}
	return direction
}

func main() {
	size := [2]int{4, 4}
	gamer := [3]int{1, 1, 0}
	m := [][]int{
		{1, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}
	result := gameDevelopment(size, gamer, m)
	fmt.Println(result)
}
