package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		words := strings.Split(input, " ")
		answer := "Case #" + strconv.Itoa(i+1) + ":"

		for i := len(words) - 1; i >= 0; i-- {
			answer += " " + words[i]
		}

		fmt.Fprintln(writer, answer)
	}
}
