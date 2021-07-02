package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b, c int

	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	sc.Scan()
	input := sc.Text()

	nums := strings.Split(input, " ")

	a, _ = strconv.Atoi(nums[0])
	b, _ = strconv.Atoi(nums[1])
	c, _ = strconv.Atoi(nums[2])

	if b >= c {
		_, err := writer.WriteString("-1\n")
		if err != nil {
			return
		}
	} else {
		count := a/(c-b) + 1
		_, err := writer.WriteString(strconv.Itoa(count) + "\n")
		if err != nil {
			return
		}
	}

}
