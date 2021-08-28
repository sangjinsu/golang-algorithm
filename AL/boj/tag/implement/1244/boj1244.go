package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func change(switches []int, n int) {
	if switches[n] == 1 {
		switches[n] = 0
	} else {
		switches[n] = 1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var switchNum int
	fmt.Fscanln(reader, &switchNum)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	inputs := strings.Split(line, " ")

	switches := make([]int, switchNum+1)
	for i, input := range inputs {
		num, _ := strconv.Atoi(input)
		switches[i+1] = num
	}

	var studentNum int
	fmt.Fscanln(reader, &studentNum)

	for i := 0; i < studentNum; i++ {
		var sex, num int
		fmt.Fscanf(reader, "%d %d\n", &sex, &num)

		if sex == 1 {
			for j := num; j < switchNum+1; j += num {
				change(switches, j)
			}
		} else {
			change(switches, num)
			j := 1
			for num-j >= 1 && num+j < switchNum+1 && switches[num-j] == switches[num+j] {
				change(switches, num-j)
				change(switches, num+j)
				j++
			}
		}
	}

	for i := 1; i < switchNum+1; i++ {
		fmt.Fprint(writer, strconv.Itoa(switches[i])+" ")
		if i%20 == 0 {
			fmt.Fprint(writer, "\n")
		}
	}
}
