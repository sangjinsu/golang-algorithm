package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscanf(reader, "%d %d\n", &n, &s)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	numStr := strings.Split(input, " ")

	var nums []int
	for _, item := range numStr {
		num, _ := strconv.Atoi(item)
		nums = append(nums, num)
	}

	var cnt int
	for i := 1; i < 1<<n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if i & (1 << j) > 0 {
				sum += nums[j]
			}
		}
		if sum == s{
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
}
