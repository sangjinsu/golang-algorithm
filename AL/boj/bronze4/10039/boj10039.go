package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var score int

	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	var sum int
	for i := 0; i < 5; i++ {
		sc.Scan()
		score, _ = strconv.Atoi(sc.Text())
		if score < 40 {
			score = 40
		}
		sum += score
	}

	result := sum / 5
	_, err := writer.WriteString(strconv.Itoa(result))
	if err != nil {
		return
	}
}
