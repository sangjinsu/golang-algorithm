package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	stack := []int{nums[0]}
	for _, num := range nums {
		if stack[len(stack)-1] < num {
			stack = append(stack, num)
		} else {
			left := 0
			right := len(stack) - 1
			for left < right {
				mid := (left + right) / 2
				if stack[mid] < num {
					left = mid + 1
				} else {
					right = mid
				}
			}
			stack[right] = num
		}
	}
	fmt.Fprint(writer, len(stack))
}
