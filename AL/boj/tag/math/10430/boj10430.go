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

	var a, b, c int
	fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)

	results := make([]int, 4)
	results[0] = (a + b) % c
	results[1] = ((a % c) + (b % c)) % c
	results[2] = (a * b) % c
	results[3] = ((a % c) * (b % c)) % c

	for _, result := range results {
		fmt.Fprintln(writer, result)
	}
}
