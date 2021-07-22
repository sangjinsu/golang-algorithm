package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func changeBinary(num int) string{
	if num <= 1 {
		return "1"
	}
	return changeBinary(num / 2) + strconv.Itoa(num % 2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer:= bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	result := changeBinary(n)
	fmt.Fprint(writer, result)
}
