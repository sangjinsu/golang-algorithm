package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	maxValue, minValue := math.MinInt64, math.MaxInt64

	for _, num := range nums {
		if maxValue < num {
			maxValue = num
		}
		if minValue > num {
			minValue = num
		}
	}

	fmt.Fprintln(writer, maxValue*minValue)
}
