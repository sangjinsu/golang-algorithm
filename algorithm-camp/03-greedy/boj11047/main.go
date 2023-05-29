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

	var N, K int

	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	coins := make([]int, N)
	for i := 0; i < N; i++ {
		var coin int
		fmt.Fscanf(reader, "%d\n", &coin)
		coins[i] = coin
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	var result int
	for _, coin := range coins {
		result += K / coin
		K %= coin
		if K == 0 {
			break
		}
	}

	fmt.Fprintf(writer, "%d", result)
}
