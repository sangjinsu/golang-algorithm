package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	var v int
	fmt.Fscan(reader, &v)

	var count int
	for _, num := range nums {
		if num == v {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}
