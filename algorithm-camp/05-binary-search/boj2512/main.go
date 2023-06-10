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

	var N, M int
	var budgets []int

	fmt.Fscanf(reader, "%d\n", &N)

	for i := 0; i < N; i++ {
		var budget int
		fmt.Fscanf(reader, "%d ", &budget)
		budgets = append(budgets, budget)
	}
	sort.Ints(budgets)

	fmt.Fscanf(reader, "%d\n", &M)

	var left, result int
	right := budgets[N-1]

	for left <= right {
		mid := (left + right) / 2
		var total int
		for _, budget := range budgets {
			if budget >= mid {
				total += mid
			} else {
				total += budget
			}
		}

		if total > M {
			right = mid - 1
		} else {
			result = mid
			left = mid + 1
		}
	}

	fmt.Fprint(writer, result)
}
