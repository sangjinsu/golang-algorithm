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

	var N, M int
	fmt.Fscanf(reader, "%d %d", &N, &M)
	result := N - M
	if result < 0 {
		result *= -1
	}
	fmt.Fprint(writer, result)
}
