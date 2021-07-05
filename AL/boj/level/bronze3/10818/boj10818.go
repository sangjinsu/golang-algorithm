package main

import (
	"bufio"
	"fmt"
	"math"
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

	var t int
	_, err := fmt.Fscanln(reader, &t)
	if err != nil {
		return
	}

	var input string
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	min, max := math.MaxInt64, math.MinInt64
	for _, n := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(n)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	_, err3 := writer.WriteString(strconv.Itoa(min) + " " + strconv.Itoa(max))
	if err3 != nil {
		return
	}
}
