package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var num1, num2 string

	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanf(reader, "%s %s", &num1, &num2)
	if err != nil {
		return
	}

	bigNum1, _ := new(big.Int).SetString(num1, 10)
	bigNum2, _ := new(big.Int).SetString(num2, 10)

	var result big.Int
	result.Add(bigNum1, bigNum2)

	fmt.Println(result.String())
}
