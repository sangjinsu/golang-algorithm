package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	M       int
	N       int
	writer  = bufio.NewWriter(os.Stdout)
	nums    []int
	visited []bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	nums = make([]int, N)
	for i := range nums {
		fmt.Fscan(reader, &nums[i])
	}

	sort.Ints(nums)

	visited = make([]bool, N)

	for i, num := range nums {
		visited[i] = true
		recursion(1, []int{num})
		visited[i] = false
	}

}

func recursion(depth int, value []int) {
	if depth == M {
		for _, v := range value {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
		return
	}

	for i := 0; i < N; i++ {
		if visited[i] == false && value[len(value)-1] <= nums[i] {
			visited[i] = true
			value = append(value, nums[i])
			recursion(depth+1, value)
			value = value[:len(value)-1]
			visited[i] = false
		}
	}
}
