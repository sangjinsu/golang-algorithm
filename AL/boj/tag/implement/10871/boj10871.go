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

	var N, X int
	fmt.Fscanf(reader, "%d %d\n", &N, &X)

	var nums []int
	for i := 0; i < N; i++ {
		var n int
		fmt.Fscanf(reader, "%d", &n)
		nums = append(nums, n)
	}

	numsLessThanN := make([]int, 0, N)
	for i := 0; i < N; i++ {
		if nums[i] < X {
			numsLessThanN = append(numsLessThanN, nums[i])
		}
	}

	strNumsLessThanN := make([]string, 0, len(numsLessThanN))
	for _, num := range numsLessThanN {
		n := strconv.Itoa(num)
		strNumsLessThanN = append(strNumsLessThanN, n)
	}
	result := strings.Join(strNumsLessThanN, " ")

	fmt.Fprint(writer, result)
}
