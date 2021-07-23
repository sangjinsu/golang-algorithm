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
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	var n, k int
	_, err := fmt.Fscan(reader, &n, &k)
	if err != nil {
		return
	}

	var queue []int
	var result []string

	for i := 0; i < n; i++ {
		queue = append(queue, i+1)
	}

	index := k - 1
	for len(queue) > 0 {
		item := queue[index]
		queue = append(queue[:index], queue[index+1:]...)
		result = append(result, strconv.Itoa(item))

		if len(queue) > 0 {
			index = (index + k - 1) % len(queue)
		}
	}

	_, err2 := fmt.Fprint(writer, "<"+strings.Join(result, ", ")+">")
	if err2 != nil {
		return
	}

}
