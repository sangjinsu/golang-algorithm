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

	thirtyNums := make([]bool, 31)

	var num int
	for i := 0; i < 28; i++ {
		fmt.Fscanln(reader, &num)
		thirtyNums[num] = true
	}

	for i := 1; i < 31; i++ {
		if !thirtyNums[i] {
			fmt.Fprintln(writer, i)
		}
	}
}
