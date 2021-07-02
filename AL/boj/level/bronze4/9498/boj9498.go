package main

import "fmt"

func main() {
	var score uint8

	_, err := fmt.Scanf("%d", &score)
	if err != nil {
		return 
	}

	result := "F"

	if score >= 90 && score <= 100 {
		result = "A"
	} else if score >= 80 && score <= 89{
		result = "B"
	} else if score >= 70 && score <= 79 {
		result = "C"
	} else if score >= 60 && score <= 69 {
		result = "D"
	}

	fmt.Println(result)
}
