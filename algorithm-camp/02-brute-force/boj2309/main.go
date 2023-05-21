package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	heights := make(map[int]int, 9)
	for i := 0; i < 9; i++ {
		var height int
		fmt.Fscanf(reader, "%d\n", &height)
		heights[i] = height
	}

	var sum int
	for _, height := range heights {
		sum += height
	}

	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if sum-(heights[i]+heights[j]) == 100 {
				delete(heights, i)
				delete(heights, j)
				break
			}
		}
		if len(heights) == 7 {
			break
		}
	}

	var heightsSlice []int
	for _, height := range heights {
		heightsSlice = append(heightsSlice, height)
	}

	sort.Ints(heightsSlice)

	for _, height := range heightsSlice {
		fmt.Fprintf(writer, "%d\n", height)
	}
}
