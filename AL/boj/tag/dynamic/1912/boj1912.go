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
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < n; i++ {
		if nums[i] < nums[i]+dp[i-1] {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		if result < dp[i] {
			result = dp[i]
		}
	}

	fmt.Fprintln(writer, result)
}
