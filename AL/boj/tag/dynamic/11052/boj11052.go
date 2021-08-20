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

	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	inputs := strings.Split(input, " ")

	cardPacks := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		cardPacks[i], _ = strconv.Atoi(inputs[i-1])
	}

	dp := make([]int, n+1)
	copy(dp, cardPacks)

	for i := 1; i < n + 1; i++ {
		for j := 1; j < i; j++ {
			comp := dp[i-j] + cardPacks[j]
			if dp[i] < comp {
				dp[i] = comp
			}
		}
	}
	fmt.Fprint(writer, strconv.Itoa(dp[n]))
}
