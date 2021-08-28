package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var dices [][]int
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		inputs := strings.Split(line, " ")
		var dice []int
		for _, input := range inputs {
			num, _ := strconv.Atoi(input)
			dice = append(dice, num)
		}
		dices = append(dices, dice)
	}
	side := map[int]int{
		0: 5, 1: 3, 2: 4, 3: 1, 4: 2, 5: 0,
	}

	maxSum := 0
	for i := 1; i < 7; i++ {
		sum := 0
		bottom := i
		for _, dice := range dices {
			bottomIndex := 0
			for j, d := range dice {
				if bottom == d {
					bottomIndex = j
				}
			}

			topIndex := side[bottomIndex]
			bottom = dice[topIndex]

			m := 0
			for j, d := range dice {
				if j != bottomIndex && j != topIndex {
					if m < d {
						m = d
					}
				}
			}

			sum += m
		}
		if maxSum < sum {
			maxSum = sum
		}
	}

	fmt.Fprint(writer, maxSum)
}
