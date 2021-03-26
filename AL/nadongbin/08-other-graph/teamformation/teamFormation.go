package main

import "fmt"

func findParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent, parent[x])
	}
	return parent[x]
}

func unionParent(parent []int, a, b int) {
	a = findParent(parent, a)
	b = findParent(parent, b)
	if a > b {
		parent[a] = b
	} else {
		parent[b] = a
	}
}

func main() {
	var N, M int
	fmt.Scanln(&N, &M)

	parent := make([]int, N+1)

	for i := 1; i < len(parent); i++ {
		parent[i] = i
	}

	for i := 0; i < M; i++ {
		var operation, a, b int
		fmt.Scanln(&operation, &a, &b)
		if operation == 0 {
			unionParent(parent, a, b)
		} else {
			if findParent(parent, a) == findParent(parent, b) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}

}
