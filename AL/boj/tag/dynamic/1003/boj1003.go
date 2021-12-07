package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibo(n int) (int, int) {
	zero := make([]int, 41)
	one := make([]int, 41)

	zero[0] = 1
	one[1] = 1

	for i := 2; i < n+1; i++ {
		zero[i] = zero[i-1] + zero[i-2]
		one[i] = one[i-1] + one[i-2]
	}

	return zero[n], one[n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscanln(reader, &N)

		one, zero := fibo(N)
		fmt.Fprintln(writer, one, zero)
	}

}
