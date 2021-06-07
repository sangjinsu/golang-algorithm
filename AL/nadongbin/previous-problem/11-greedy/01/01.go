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
	var N int
	fmt.Scanf("%d\n", &N)

	var input string
	// r := bufio.NewReader(os.Stdin)
	// input, _ = r.ReadString('\n')
	// input = strings.TrimSuffix(input, "\r\n")
	// fmt.Printf("%#v", input)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input = s.Text()
	fmt.Printf("%#v", input)

	inputStr := strings.Split(input, " ")

	adventurers := make([]int, len(inputStr))
	for i := 0; i < len(inputStr); i++ {
		adventurers[i], _ = strconv.Atoi(inputStr[i])
	}

	sort.Slice(adventurers, func(i, j int) bool {
		return adventurers[i] < adventurers[j]
	})

	group := 0
	count := 0
	fmt.Println(adventurers)
	for _, v := range adventurers {
		count++
		if count >= v {
			group++
			count = 0
		}
	}

	fmt.Println(group)
}
