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

	var number int
	fmt.Fscanln(reader, &number)

	var k int
	fmt.Fscanln(reader, &k)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	brokens := strings.Split(line, " ")

	result := 100 - number
	if result < 0 {
		result *= -1
	}
	for i := 0; i < 1_000_000; i++ {
		num := strconv.Itoa(i)
		isBroken := false
		for j := 0; j < len(num); j++ {
			for _, broken := range brokens {
				if string(num[j]) == broken {
					isBroken = true
					break
				}
			}
			if isBroken {
				break
			}
		}
		if isBroken == false {
			cnt := i - number
			if cnt < 0 {
				cnt *= -1
			}
			if result > len(num)+cnt {
				result = len(num) + cnt
			}

		}
	}
	fmt.Fprint(writer, result)
}
