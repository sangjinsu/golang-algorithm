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
	var n int
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	inputArr := strings.Split(input, " ")

	soldiers := make([]int, len(inputArr))

	for i, v := range inputArr {
		soldiers[i], _ = strconv.Atoi(v)
	}

	dp := make([]int, len(soldiers))

	for i := len(soldiers) - 1; i >= 0; i-- {
		for j := i; j < len(soldiers); j++ {
			if soldiers[j] < soldiers[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	fmt.Println(dp)
	max := -1
	for _, v := range dp {
		if max < v {
			max = v
		}
	}

	fmt.Println(n - (max + 1))
}
