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

	var A, B, C int
	fmt.Fscanf(reader, "%d %d %d\n", &A, &B, &C)
	result := 1
	for i := 0; i < B; i++ {
		result *= A
	}
	result %= C

	fmt.Fprintln(writer, result)
}
