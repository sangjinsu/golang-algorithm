package main

func dfs1(v int, visited []bool, graph [][]int) {
	stack := []int{v}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[node] = true
		// fmt.Println(node)
		for _, nd := range graph[node] {
			if visited[nd] == false {
				stack = append(stack, nd)
			}
		}
	}
}

func dfs2(v int, visited []bool, graph [][]int) {
	visited[v] = true
	// fmt.Println(v)
	for _, node := range graph[v] {
		if visited[node] == false {
			dfs2(node, visited, graph)
		}
	}
}

func main() {
	graph := [][]int{
		{},
		{2, 6, 5},
		{1},
		{4, 5, 6},
		{3},
		{3},
		{1, 3},
	}
	n := 6
	visited := make([]bool, n+1)
	dfs1(1, visited, graph)

	visited = make([]bool, n+1)
	dfs2(1, visited, graph)
}
