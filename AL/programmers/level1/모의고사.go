package level1

func solution(answers []int) []int {
	supoza1 := []int{1, 2, 3, 4, 5}
	supoza2 := []int{2, 1, 2, 3, 2, 4, 2, 5}
	supoza3 := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	scores := []int{0, 0, 0}

	for i, v := range answers {

		if supoza1[i%5] == v {
			scores[0]++
		}

		if supoza2[i%8] == v {
			scores[1]++
		}

		if supoza3[i%10] == v {
			scores[2]++
		}
	}

	var max int
	for _, score := range scores{
		if max < score {
			max = score
		}
	}

	var result []int
	for i, score := range scores {
		if max == score {
			result = append(result, i + 1)
		}
	}

	return result
}
