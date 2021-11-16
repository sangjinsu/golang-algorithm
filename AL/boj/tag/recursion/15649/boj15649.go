package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func recursion(cnt int, result []int, visited []bool, N int, M int) {
	if cnt == M {
		var strResult string
		for _, n := range result {
			strResult += strconv.Itoa(n) + " "
		}
		fmt.Fprintln(writer, strResult)
		return
	}
	for i := 1; i < N+1; i++ {
		if visited[i] == false {
			visited[i] = true
			result = append(result, i)
			recursion(cnt+1, result, visited, N, M)
			visited[i] = false
			result = result[:len(result)-1]
		}
	}
}

func main() {

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	result := make([]int, 0, M)
	visited := make([]bool, N+1)
	recursion(0, result, visited, N, M)
}
