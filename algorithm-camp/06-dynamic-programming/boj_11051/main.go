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

	var N, K int
	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	results := make([][]int, 1001)
	for i := 0; i < 1001; i++ {
		results[i] = make([]int, 1001)
	}

	for i := 0; i < 1001; i++ {
		results[i][0], results[i][i] = 1, 1
		for j := 1; j < i; j++ {
			results[i][j] = results[i-1][j-1] + results[i-1][j]
			results[i][j] %= 10_007
		}
	}

	fmt.Fprintln(writer, results[N][K]%10_007)

}
