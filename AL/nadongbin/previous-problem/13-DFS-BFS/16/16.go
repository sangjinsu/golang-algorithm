package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	fmt.Scanln(&N, &M)

	data := make([][]int, N)
	temp := make([][]int, N)

	ro := make([]int, M)
	for i := 0; i < N; i++ {
		row := make([]int, len(ro))
		copy(row, ro)
		temp[i] = row
	}

	r := bufio.NewReader(os.Stdin)

	for i := 0; i < N; i++ {
		input, _ := r.ReadString('\n')
		input = strings.TrimSuffix(input, "\r\n")
		inputs := strings.Split(input, " ")

		for _, v := range inputs {
			value, _ := strconv.Atoi(v)
			data[i] = append(data[i], value)
		}
	}

	result := dfs(data, temp, 0, N, M)
	fmt.Println(result)
}

func virus(x, y int, N, M int, temp [][]int) {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := x + dy[i]
		if nx >= 0 && nx < N && ny >= 0 && ny < M {
			if temp[nx][ny] == 0 {
				temp[nx][ny] = 2
				virus(nx, ny, N, M, temp)
			}
		}
	}
}

func getScore(temp [][]int) int {
	score := 0
	for _, t := range temp {
		for _, v := range t {
			if v == 0 {
				score++
			}
		}
	}
	return score
}

func dfs(data, temp [][]int, count int, N, M int) int {
	result := 0
	if count == 3 {
		for i, d := range data {
			for j, v := range d {
				temp[i][j] = v
			}
		}
		for i, t := range temp {
			for j, v := range t {
				if v == 2 {
					virus(i, j, N, M, temp)
				}
			}
		}
		result = int(math.Max(float64(result), float64(getScore(temp))))
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if data[i][j] == 0 {
				data[i][j] = 1
				count++
				dfs(data, temp, count, N, M)
				data[i][j] = 0
				count--
			}
		}
	}

	return result
}
