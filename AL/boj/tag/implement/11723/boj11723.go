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

	var M int
	fmt.Fscanf(reader, "%d\n", &M)

	numSet := make(map[int]int, 20)
	for i := 0; i < M; i++ {
		var command string
		var num int
		fmt.Fscanf(reader, "%s %d\n", &command, &num)
		switch command {
		case "add":
			if _, ok := numSet[num]; !ok {
				numSet[num] = num
			}
		case "remove":
			if _, ok := numSet[num]; ok {
				delete(numSet, num)
			}
		case "check":
			if _, ok := numSet[num]; !ok {
				fmt.Fprintf(writer, "0\n")
			} else {
				fmt.Fprintf(writer, "1\n")
			}
		case "toggle":
			if _, ok := numSet[num]; !ok {
				numSet[num] = num
			} else {
				delete(numSet, num)
			}
		case "all":
			for value := 1; value < 21; value++ {
				if _, ok := numSet[value]; !ok {
					numSet[value] = value
				}
			}
		case "empty":
			numSet = make(map[int]int, 20)
		}

	}
}
