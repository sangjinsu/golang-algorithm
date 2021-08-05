package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getPrimes() []bool {
	var primes []bool
	n := 1000+1
	for i := 0; i < n; i++ {
		primes = append(primes, true)
	}

	primes[0], primes[1] = false, false

	m := int(math.Sqrt(float64(n)))

	for i := 2; i < m+1; i++ {
		if primes[m] {
			for j := i + i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	return primes
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	inputs := strings.Split(input, " ")

	var nums []int
	for _, s := range inputs {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}

	cnt := 0
	primes := getPrimes()

	for i := 0; i < n; i++ {
		if primes[nums[i]] {
			cnt++
		}
	}

	fmt.Fprint(writer, strconv.Itoa(cnt))

}
