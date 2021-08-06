package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func makePrimes() [1000001]bool {
	k := 1000000
	var primes [1000001]bool

	primes[0], primes[1] = true, true

	m := int(math.Sqrt(float64(k)))

	for i := 2; i < m+1; i++ {
		if !primes[i] {
			for j := i * 2; j < k+1; j += i {
				primes[j] = true
			}
		}
	}

	return primes
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	primes := makePrimes()
	for true {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		num, _ := strconv.Atoi(input)

		if num == 0 {
			return
		}

		result := "Goldbach's conjecture is wrong."
		for i := 2; i < len(primes); i++ {
			if primes[i] == false && primes[num-i] == false {
				result = fmt.Sprintf("%d = %d + %d", num, i, num-i)
				break
			}
		}

		fmt.Fprintln(writer, result)
	}
}
