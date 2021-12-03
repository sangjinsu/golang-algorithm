package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	recursion([]string{}, N, M)
}

func recursion(nums []string, N, M int) {
	if len(nums) == M {
		result := strings.Join(nums, " ")
		fmt.Fprintln(writer, result)
		return
	}

	for i := 1; i < N+1; i++ {
		nums = append(nums, strconv.Itoa(i))
		recursion(nums, N, M)
		nums = nums[:len(nums)-1]
	}
}
