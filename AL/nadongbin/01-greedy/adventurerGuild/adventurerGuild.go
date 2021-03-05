package main

import (
	"fmt"
	"sort"
)

func adventurerGuild(adventurers []int) int {
	sort.Ints(adventurers)

	count := 0
	guild := 0
	for _, adventurer := range adventurers {
		count++
		if count >= adventurer {
			guild++
			count = 0
		}
	}

	return guild
}

func main() {
	fmt.Println(adventurerGuild([]int{2, 3, 1, 2, 2}))
}
