package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	visited []bool
	values  []int
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	defer writer.Flush()

	var N, M int
	fmt.Fscanln(reader, &N, &M)

	values = make([]int, N)
	for i := range values {
		fmt.Fscan(reader, &values[i])
	}

	sort.Ints(values)

	visited = make([]bool, N)
	recursion([]string{}, N, M)
}

func recursion(nums []string, N, M int) {
	if len(nums) == M {
		result := strings.Join(nums, " ")
		fmt.Fprintln(writer, result)
		return
	}

	for i := 0; i < N; i++ {
		if visited[i] == false {
			visited[i] = true
			nums = append(nums, strconv.Itoa(values[i]))
			recursion(nums, N, M)
			nums = nums[:len(nums)-1]
			visited[i] = false
		}
	}
}
