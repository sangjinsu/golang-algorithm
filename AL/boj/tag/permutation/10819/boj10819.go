package main

import (
	"bufio"
	"fmt"
	"os"
)

var writer = bufio.NewWriter(os.Stdout)
var visited []bool
var nums []int
var result int
var N int

func main() {
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d\n", &N)

	nums = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	visited = make([]bool, N)
	recursion(0, []int{})
	fmt.Fprintln(writer, result)
}

func recursion(idx int, ints []int) {
	if idx == N {
		sum := 0
		for i := 0; i < N-1; i++ {
			v := ints[i] - ints[i+1]
			if v < 0 {
				v *= -1
			}
			sum += v
		}

		if result < sum {
			result = sum
		}
	}

	for i := 0; i < N; i++ {
		if visited[i] == false {
			visited[i] = true
			ints = append(ints, nums[i])
			recursion(idx+1, ints)
			visited[i] = false
			ints = ints[:len(ints)-1]
		}
	}
}
