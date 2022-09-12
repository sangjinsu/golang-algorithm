package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var M, N int
	fmt.Fscanf(reader, "%d %d\n", &M, &N)

	monsters := make(map[string]int)
	names := make(map[int]string)

	for i := 1; i <= M; i++ {
		var name string
		fmt.Fscanln(reader, &name)
		monsters[name] = i
		names[i] = name
	}

	for i := 0; i < N; i++ {
		var value string
		fmt.Fscanln(reader, &value)
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Fprintln(writer, monsters[value])
		} else {
			fmt.Fprintln(writer, names[number])
		}
	}
}
