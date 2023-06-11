package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int

	fmt.Fscanf(reader, "%d\n", &N)
	cards := make([]int, N)
	for i := 0; i < N; i++ {
		var card int
		fmt.Fscanf(reader, "%d ", &card)
		cards[i] = card
	}

	fmt.Fscanf(reader, "%d\n", &M)
	nums := make([]int, M)
	for i := 0; i < M; i++ {
		var num int
		fmt.Fscanf(reader, "%d ", &num)
		nums[i] = num
	}

	indicies := make(map[int]int, N)
	for _, card := range cards {
		indicies[card] = 1
	}

	for _, num := range nums {
		if _, ok := indicies[num]; ok {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}

}
