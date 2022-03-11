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

	var GCD int
	if A >= B {
		GCD = gcd(B, A)
	} else {
		GCD = gcd(A, B)
	}

	result := A * B / GCD

	fmt.Fprintln(writer, result)
}

func gcd(b int, a int) int {
	for a != 0 {
		b, a = a, b%a
	}
	return b
}
