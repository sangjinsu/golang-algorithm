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

	var number string
	fmt.Fscan(reader, &number)
	cnt, answer := isMultipleOfThree(number, 0)

	fmt.Fprintln(writer, cnt)
	fmt.Fprint(writer, answer)

}

func isMultipleOfThree(number string, count int) (int, string) {
	if len(number) == 1 {
		n, _ := strconv.Atoi(number)
		if n%3 == 0 {
			return count, "YES"
		} else {
			return count, "NO"
		}
	}
	count++
	nums := strings.Split(number, "")
	total := 0
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		total += n
	}
	t := strconv.Itoa(total)
	return isMultipleOfThree(t, count)
}
