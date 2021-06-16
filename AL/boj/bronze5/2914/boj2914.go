package main

import "fmt"

func main() {
	var songs, average uint

	_, err := fmt.Scanf("%d %d", &songs, &average)
	if err != nil {
		return
	}
	melodies := songs*(average-1) + 1

	fmt.Println(melodies)
}
