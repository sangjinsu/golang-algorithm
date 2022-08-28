package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscanln(reader, &T)

	var R int
	var S string
	for i := 0; i < T; i++ {
		fmt.Fscanf(reader, "%d %s\n", &R, &S)
		var result string
		for _, ch := range S {
			result += strings.Repeat(string(ch), R)
		}
		fmt.Fprintln(writer, result)
	}
}
