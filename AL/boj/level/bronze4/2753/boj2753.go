package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	input := sc.Text()
	year, _ := strconv.Atoi(input)

	if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
		fmt.Printf("%d", 1)
	} else {
		fmt.Printf("%d", 0)
	}
}
