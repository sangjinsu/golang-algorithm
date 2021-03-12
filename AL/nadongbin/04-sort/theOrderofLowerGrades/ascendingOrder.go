package main

import (
	"fmt"
	"sort"
)

type student struct {
	name  string
	score int
}

func main() {

	var n int
	fmt.Scanln(&n)

	arr := make([]student, n)

	for i := 0; i < n; i++ {
		var name string
		var score int
		fmt.Scanln(&name, &score)
		arr[i] = student{name: name, score: score}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score < arr[j].score
	})

	for _, student := range arr {
		fmt.Printf("%s ", student.name)
	}
}
