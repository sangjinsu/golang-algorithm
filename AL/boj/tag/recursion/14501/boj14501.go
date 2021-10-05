package main

import (
	"bufio"
	"fmt"
	"os"
)

var result = 0

func recursion(T, P []int, day, N, value int) {
	if day >= N {
		if result < value {
			result = value
		}
		return
	}
	recursion(T, P, day+1, N, value)
	if day+T[day] <= N {
		recursion(T, P, day+T[day], N, value+P[day])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	T := make([]int, N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		var t, p int
		fmt.Fscanf(reader, "%d %d\n", &t, &p)
		T[i] = t
		P[i] = p
	}
	recursion(T, P, 0, N, 0)
	fmt.Fprint(writer, result)
}
