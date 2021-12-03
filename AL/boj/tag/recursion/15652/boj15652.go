package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var writer = bufio.NewWriter(os.Stdout)

func main() {
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	recursion(0, N, M, []string{})
}

func recursion(idx, n, m int, value []string) {
	if idx == m {
		fmt.Fprintln(writer, strings.Join(value, " "))
		return
	}

	var start = 1
	if len(value) > 0 {
		start, _ = strconv.Atoi(value[len(value)-1])
	}

	for i := start; i < n+1; i++ {
		value = append(value, strconv.Itoa(i))
		recursion(idx+1, n, m, value)
		value = value[:len(value)-1]
	}
}
