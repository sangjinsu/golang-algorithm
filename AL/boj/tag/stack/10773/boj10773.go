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

	var k int
	var stack []int
	fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		var num int
		fmt.Fscan(reader, &num)

		if num == 0 {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, num)
		}
	}

	var sum int
	for _, num := range stack {
		sum += num
	}

	fmt.Fprint(writer, sum)
}
