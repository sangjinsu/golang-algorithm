package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")
	inputStr := strings.Split(input, " ")

	loc := make([]int, len(inputStr))
	for i, s := range inputStr {
		loc[i], _ = strconv.Atoi(s)
	}

	sort.Ints(loc)

	result := loc[len(loc)/2-1]

	fmt.Println(result)
}
