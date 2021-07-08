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
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {

		}
	}(writer)

	nums := make([]int, 3)
	_, err := fmt.Fscan(reader, &nums[0], &nums[1], &nums[2])
	if err != nil {
		return
	}

	sort.Ints(nums)
	middle := nums[1]

	_, err = fmt.Fprintln(writer, middle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
}
