package main

import "testing"

const (
	n = 6
)

var graph = [][]int{
	{},
	{2, 6, 5},
	{1},
	{4, 5, 6},
	{3},
	{3},
	{1, 3},
}

func BenchmarkDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		visited := make([]bool, n+1)
		dfs1(1, visited, graph)
	}
}

func BenchmarkDFSRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		visited := make([]bool, n+1)
		dfs2(1, visited, graph)
	}
}
