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
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	inputs := strings.Split(line, " ")

	nums := make([]int, n)
	for i, input := range inputs {
		num, _ := strconv.Atoi(input)
		nums[i] = num
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}

		}

		if result < dp[i] {
			result = dp[i]
		}
	}

	fmt.Fprint(writer, result)
}
