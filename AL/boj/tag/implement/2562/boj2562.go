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

	var num int
	var maxNum int
	var maxIndex int
	for i := 1; i < 10; i++ {
		fmt.Fscanln(reader, &num)
		if num > maxNum {
			maxNum = num
			maxIndex = i
		}
	}

	fmt.Fprintln(writer, maxNum)
	fmt.Fprint(writer, maxIndex)
}
