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

	var n int
	fmt.Fscanln(reader, &n)

	var maxNums []int
	var maxCnt int
	for i := n / 2; i < n+1; i++ {
		var nums []int
		cnt := 0
		first := n
		second := i
		for first >= 0 {
			nums = append(nums, first)
			first, second = second, first-second
			cnt++
		}

		if maxCnt < cnt {
			maxCnt = cnt
			maxNums = nums
		}
	}

	var result []string
	for _, num := range maxNums {
		result = append(result, strconv.Itoa(num))
	}
	fmt.Fprintln(writer, maxCnt)
	fmt.Fprint(writer, strings.Join(result, " "))
}
