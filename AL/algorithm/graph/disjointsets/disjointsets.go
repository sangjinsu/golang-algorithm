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

	for i := 0; i < e; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		unionParent(parent, a, b)
	}

	fmt.Fprint(writer, "각 원소가 속한 집합: ")
	for i := 1; i < v+1; i++ {
		fmt.Fprint(writer, findParent(parent, i))
		fmt.Fprint(writer, " ")
	}
	fmt.Fprintln(writer)

	fmt.Fprint(writer, "부모 테이블: ")
	for i := 1; i < v+1; i++ {
		fmt.Fprint(writer, parent[i])
		fmt.Fprint(writer, " ")
	}
}
