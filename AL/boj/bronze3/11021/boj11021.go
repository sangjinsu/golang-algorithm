package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	_, err := fmt.Fscan(reader, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var a, b int
		_, err2 := fmt.Fscan(reader, &a, &b)
		if err2 != nil {
			return
		}
		_, err3 := writer.WriteString(
			"Case #" + strconv.Itoa(i + 1) + ": " + strconv.Itoa(a+b) + "\n")
		if err3 != nil {
			return
		}
	}
}
