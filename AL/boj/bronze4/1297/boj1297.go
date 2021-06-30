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
	var diagonal, heightRatio, widthRatio int

	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	sc.Scan()
	input := strings.Split(sc.Text(), " ")

	diagonal, _ = strconv.Atoi(input[0])
	heightRatio, _ = strconv.Atoi(input[1])
	widthRatio, _ = strconv.Atoi(input[2])

	x := math.Sqrt(
		math.Pow(float64(diagonal), 2) / (math.Pow(float64(heightRatio), 2) + math.Pow(float64(widthRatio), 2)))

	height := strconv.Itoa(int(float64(heightRatio) * x))
	width := strconv.Itoa(int(float64(widthRatio) * x))

	_, err := writer.WriteString(height + " " + width)
	if err != nil {
		fmt.Println("??")
	}

}
