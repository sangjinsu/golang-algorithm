package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var teams [][]int

func makeTeam(idx int, team []int, visited []bool, N int) {
	if len(team) == N/2 {
		copySlice := make([]int, len(team))
		copy(copySlice, team)
		teams = append(teams, copySlice)
		return
	}
	for i := idx; i < N; i++ {
		if visited[i] == false {
			visited[i] = true
			team = append(team, i)
			makeTeam(i, team, visited, N)
			team = team[:len(team)-1]
			visited[i] = false
		}
	}
}

func calScore(teamA []int, teamB []int, mat [][]int) int {
	var scoreA, scoreB int
	for i := 0; i < len(teamA)-1; i++ {
		for j := i + 1; j < len(teamA); j++ {
			scoreA += mat[teamA[i]][teamA[j]] + mat[teamA[j]][teamA[i]]
		}
	}
	for i := 0; i < len(teamB)-1; i++ {
		for j := i + 1; j < len(teamB); j++ {
			scoreB += mat[teamB[i]][teamB[j]] + mat[teamB[j]][teamB[i]]
		}
	}

	result := scoreA - scoreB

	if result >= 0 {
		return result
	}
	return -1 * result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)
	var mat [][]int

	for i := 0; i < N; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		line := strings.Split(input, " ")

		nums := make([]int, len(line))
		for j, char := range line {
			num, _ := strconv.Atoi(char)
			nums[j] = num
		}
		mat = append(mat, nums)
	}

	visited := make([]bool, N)
	makeTeam(0, []int{}, visited, N)
	teamA := teams[len(teams)/2:]
	var teamB [][]int
	temp := teams[:len(teams)/2]
	for i := len(temp) - 1; i >= 0; i-- {
		teamB = append(teamB, temp[i])
	}

	result := 9999
	for i := 0; i < len(teamA); i++ {
		value := calScore(teamA[i], teamB[i], mat)
		if result > value {
			result = value
		}
	}

	fmt.Fprint(writer, result)
}
