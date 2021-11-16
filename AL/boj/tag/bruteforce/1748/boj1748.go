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

	var result int
	for i := 1; i <= N; i *= 10 {
		result += N - i + 1
	}

	fmt.Fprintln(writer, result)
}
