package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	inputArr := strings.Split(input, " ")
	warehouse := make([]int, n)

	var err error
	for i, food := range inputArr {
		warehouse[i], err = strconv.Atoi(food)
		if err != nil {
			log.Fatal("Can not convert string to int")
		}
	}

	dp := make([]int, n)
	dp[0] = warehouse[0]
	dp[1] = int(math.Max(float64(warehouse[0]), float64(warehouse[1])))

	for i := 2; i < n; i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+warehouse[i]), float64(dp[i-1])))
	}

	result := dp[len(dp)-1]
	fmt.Println(result)
}
