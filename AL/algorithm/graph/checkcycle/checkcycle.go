package main

import (
	"bufio"
	"fmt"
	"os"
)

func findParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent, parent[x])
	}
	return parent[x]
}

func unionParent(parent []int, a, b int) {
	a = findParent(parent, a)
	b = findParent(parent, b)
	if a < b {
		parent[b] = a
	} else {
		parent[a] = b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var v, e int
	fmt.Fscanf(reader, "%d %d\n", &v, &e)

	parent := make([]int, v+1)
	for i := 1; i < v+1; i++ {
		parent[i] = i
	}

	cycle := false

	for i := 0; i < e; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		if findParent(parent, a) == findParent(parent, b) {
			cycle = true
			break
		} else {
			unionParent(parent, a, b)
		}
	}

	if cycle {
		fmt.Fprintln(writer, "Cycle")
	} else {
		fmt.Fprintln(writer, "Not Cycle")
	}
}
