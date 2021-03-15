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
	var tc int
	fmt.Scanln(&tc)

	for tc > 0 {
		var n, m int
		fmt.Scanln(&n, &m)

		dp := make([][]int, n)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		inputArr := strings.Split(input, " ")

		mine := make([]int, m*n)
		for i := 0; i < len(inputArr); i++ {
			mine[i], _ = strconv.Atoi(inputArr[i])
		}

		index := 0
		for i := 0; i < n; i++ {
			dp[i] = mine[index : index+m]
			index += m
		}

		for j := 1; j < m; j++ {
			var left, leftUp, leftDown int
			for i := 0; i < n; i++ {
				left = dp[i][j-1]

				if i != 0 {
					leftUp = dp[i-1][j-1]
				}

				if i != n-1 {
					leftDown = dp[i+1][j-1]
				}

				dp[i][j] += int(math.Max(math.Max(float64(left), float64(leftUp)), float64(leftDown)))
			}
			fmt.Println(dp)
		}
		tc--

		result := 0
		for _, row := range dp {
			result = row[len(row)-1]
		}
		fmt.Println(result)
	}
}
