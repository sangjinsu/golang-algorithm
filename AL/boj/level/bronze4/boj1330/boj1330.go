package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var a ,b int
	//
	//_, err := fmt.Scanf("%d %d", &a, &b)
	//if err != nil {
	//	return
	//}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	strArr := strings.Fields(input)
	a, _ := strconv.Atoi(strArr[0])
	b, _ := strconv.Atoi(strArr[1])

	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}
