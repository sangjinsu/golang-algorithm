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

	var A, B int
	fmt.Fscanf(reader, "%d %d\n", &A, &B)

	result := A*A - B*B
	fmt.Fprintln(writer, result)
}
