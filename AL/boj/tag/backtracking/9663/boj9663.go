package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func nQeens(k int, n int, rows []int) int {
	if k == n {
		return 1
	}
	cnt := 0
	for x := 0; x < n; x++ {
		ko := 1
		for y := 0; y < k; y++ {
			if rows[y] == x || k-y == int(math.Abs(float64(rows[y]-x))) {
				ko = 0
				break
			}
		}
		if ko == 1 {
			rows[k] = x
			cnt += nQeens(k + 1, n, rows)
		}
	}

	return cnt
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	rows := make([]int, n)
	result := nQeens(0, n, rows)

	fmt.Fprint(writer, result)
}
