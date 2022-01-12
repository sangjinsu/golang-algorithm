package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	check [1501][1501]bool
	total int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var x, y, z int
	fmt.Fscanf(reader, "%d %d %d\n", &x, &y, &z)
	total = x + y + z

	if total%3 != 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	recursion(x, y)
	if check[total/3][total/3] {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}

func recursion(x, y int) {
	if check[x][y] {
		return
	}
	check[x][y] = true
	a := [3]int{x, y, total - x - y}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i] < a[j] {
				b := [3]int{x, y, total - x - y}
				b[i] += a[i]
				b[j] -= a[i]
				recursion(b[0], b[1])
			}
		}
	}
}
