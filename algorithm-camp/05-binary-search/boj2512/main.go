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
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	var inputStr string
	var budgets []int

	fmt.Fscanf(reader, "%d\n", &N)
	fmt.Fscanf(reader, "%s\n", &inputStr)
	budgetsStr := strings.Split(inputStr, " ")
	for _, b := range budgetsStr {
		budget, _ := strconv.Atoi(b)
		budgets = append(budgets, budget)
	}
	sort.Ints(budgets)

	left := 0
	right := budgets[len(budgets)-1]

	for left < right {

	}
}
