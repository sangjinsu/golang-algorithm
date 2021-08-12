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
	defer writer.Flush()

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	nums := strings.Split(input, " ")

	m, _ := strconv.Atoi(nums[0])
	n, _ := strconv.Atoi(nums[1])

	var primes []bool
	for i := 0; i < n+1; i++ {
		primes = append(primes, true)
	}

	primes[0] = false
	primes[1] = false

	k := int(math.Sqrt(float64(n)))
	for i := 2; i < k+1; i++ {
		if primes[i] {
			for j := i * 2; j < n+1; j += i {
				primes[j] = false
			}
		}
	}

	for i := m; i <= n; i++ {
		if primes[i] {
			fmt.Fprintln(writer, strconv.Itoa(i))
		}
	}

}
