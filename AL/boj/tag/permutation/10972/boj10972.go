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

	find := false

	var N int

	fmt.Fscanf(reader, "%d\n", &N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	for i := N - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			for j := N - 1; j >= 1; j-- {
				if nums[i-1] < nums[j] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					sort.Ints(nums[i:])
					find = true
					break
				}
			}

			if find {
				for _, v := range nums {
					fmt.Fprintf(writer, "%d ", v)
				}
				fmt.Fprintln(writer)
				break
			}
		}
	}

	if find == false {
		fmt.Fprintln(writer, -1)
	}
}
