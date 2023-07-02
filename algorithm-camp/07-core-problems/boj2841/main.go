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

	var N, P int
	fmt.Fscanf(reader, "%d %d\n", &N, &P)

	var result int
	stack := make([][]int, N+1)
	for i := 0; i < N; i++ {
		var line, flat int
		fmt.Fscanf(reader, "%d %d\n", &line, &flat)

		for len(stack[line]) > 0 && stack[line][len(stack[line])-1] > flat {
			stack[line] = stack[line][:len(stack[line])-1]
			result++
		}

		if len(stack[line]) > 0 && stack[line][len(stack[line])-1] == flat {
			continue
		}

		stack[line] = append(stack[line], flat)
		result++
	}

	fmt.Fprint(writer, result)
}
