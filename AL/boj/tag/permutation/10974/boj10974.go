package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	result  []string
	visited []bool
	N       int
	writer  = bufio.NewWriter(os.Stdout)
)

func recursion(k int) {
	if k == N {
		value := strings.Join(result, " ")
		fmt.Fprintln(writer, value)
	}

	for i := 1; i < N+1; i++ {
		if visited[i] == false {
			visited[i] = true
			num := strconv.Itoa(i)
			result = append(result, num)
			recursion(k + 1)
			visited[i] = false
			result = result[:len(result)-1]
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	visited = make([]bool, N+1)

	recursion(0)
}
