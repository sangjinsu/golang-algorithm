package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	sort.Ints(nums)

	for i := 1; i < N; i++ {
		nums[i] += nums[i-1]
	}

	var result int
	for i := 0; i < N; i++ {
		result += nums[i]
	}

	fmt.Fprintln(writer, result)
}
