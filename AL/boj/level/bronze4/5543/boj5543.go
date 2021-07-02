package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	burger := 2001
	cola := 2001

	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < 3; i++ {
		sc.Scan()
		burger = compareValue(burger, sc.Text())
	}

	for i := 0; i < 2; i++ {
		sc.Scan()
		cola = compareValue(cola, sc.Text())
	}

	setMenu := burger + cola - 50
	fmt.Println(setMenu)
	_, err := writer.WriteString(strconv.Itoa(setMenu))
	if err != nil {
		return
	}
}

func compareValue(value int, comparison string) int {
	comparedValue, _ := strconv.Atoi(comparison)
	if value > comparedValue {
		return comparedValue
	}
	return value
}

