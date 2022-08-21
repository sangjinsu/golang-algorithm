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

	var word string
	fmt.Fscan(reader, &word)

	length := int('z' - 'a' + 1)
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = "-1"
	}

	for i, r := range word {
		if result[r-'a'] == "-1" {
			result[r-'a'] = strconv.Itoa(i)
		}
	}

	fmt.Fprintln(writer, strings.Join(result, " "))
}
