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

	var n int
	fmt.Fscanln(reader, &n)

	var queue []int
	for i := 1; i < n+1; i++ {
		queue = append(queue, i)
	}

	for len(queue) > 1 {
		queue = append(queue[2:], queue[1])
	}

	fmt.Fprint(writer, queue[0])
}
