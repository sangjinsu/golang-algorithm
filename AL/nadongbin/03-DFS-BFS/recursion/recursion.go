package main

import (
	"fmt"
)

func recursion(n int) {
	if n == 0 {
		fmt.Println("종료")
		return
	}
	fmt.Println("재귀 함수를 호출합니다.", n)
	recursion(n - 1)
	fmt.Println("재귀 함수를 종료합니다.", n)
}

func main() {
	recursion(100)
}
