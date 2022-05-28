package main

import (
	"bufio"
	"fmt"
	"os"
)

func recursion(A, B, C int) int {
	if B == 0 {
		return 1
	}
	if B == 1 {
		return A % C
	}

	temp := recursion(A, B/2, C)

	if B%2 == 0 {
		return temp * temp % C
	} else {
		return temp * temp % C * A % C
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var A, B, C int
	fmt.Fscanf(reader, "%d %d %d\n", &A, &B, &C)

	fmt.Fprint(writer, recursion(A, B, C))
}
