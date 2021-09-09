package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	nums := make([]int, 9)
	var sum int
	for i := 0; i < 9; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		nums[i] = n
		sum += n
	}

	sort.Ints(nums)

	idx1 := -1
	idx2 := -1
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if 100 == sum-(nums[i]+nums[j]) {
				idx1 = i
				idx2 = j
				break
			}
		}
		if idx1 >= 0 {
			break
		}
	}

	for i, num := range nums {
		if i != idx1 && i != idx2 {
			fmt.Fprintln(writer, num)
		}
	}
}
